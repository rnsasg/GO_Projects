package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	go produce(ctx)
	consume(ctx)
}

// Producer

const (
	topic          = "my-kafka-topic"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
		// wait until we get 10 messages before writing
		BatchSize: 10,
		// no matter what happens, write all pending messages
		// every 2 seconds
		BatchTimeout: 2 * time.Second,
		// can be set to -1, 0, or 1
		// 1 is a good default for most non-transactional data
		RequiredAcks: 1,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}

type Logger interface {
	Printf(string, ...interface{})
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker1Address, broker2Address, broker3Address},
		Topic:    topic,
		GroupID:  "my-group",
		Logger:   l,
		MinBytes: 5,
		// the kafka library requires you to set the MaxBytes
		// in case the MinBytes are set
		MaxBytes: 1e6,
		// wait for at most 3 seconds before receiving new data
		MaxWait: 3 * time.Second,
		// this will start consuming messages from the earliest available
		StartOffset: kafka.FirstOffset,
		// if you set it to `kafka.LastOffset` it will only consume new messages
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}
