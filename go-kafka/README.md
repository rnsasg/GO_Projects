
# Implementing a Kafka Producer and Consumer In Golang For Production

## Introduction 

Kafka is an open-source event streaming platform, used for publishing and processing events at high-throughput. There are a lot of popular libraries for Go in order to interface with Kafka. For this post, we will be using the kafka-go library (but the same concepts will apply for any other library as well).

## System Architecture 

There are a bunch of processes that we need to start to run our cluster :

Zookeeper : Which is used by Kafka to maintain state between the nodes of the cluster.
Kafka brokers : The “pipes” in our pipeline, which store and emit data.
Producers : That insert data into the cluster.
Consumers : That read data from the cluster.

<picture>
    <img alt="Kafka System Architecture" src=architecture.png>
</picture>

## Installation 

1. Install Java [link](https://www.oracle.com/java/technologies/downloads/)
2. Download Kafka tar file [link]((https://archive.apache.org/dist/kafka/3.6.0/kafka_2.13-3.6.0.tgz))
3. Untar Kafka
```
tar -xvzf kafka_2.13-3.6.0.tgz
```
4. Start ZooKeeper
```
bin/zookeeper-server-start.sh config/zookeeper.properties
```
5. Configure and Start Kafka Brokers 

```
cp config/server.properties config/server.1.properties
cp config/server.properties config/server.2.properties
cp config/server.properties config/server.3.properties
```

server.1.properties
```
broker.id=1
listeners=PLAINTEXT://:9093
log.dirs=/tmp/kafka-logs1
```

server.2.properties
```
broker.id=2
listeners=PLAINTEXT://:9094
log.dirs=/tmp/kafka-logs2
```

server.3.properties
```
broker.id=3
listeners=PLAINTEXT://:9095
log.dirs=/tmp/kafka-logs3
```

Start each one in different terminals
```
bin/kafka-server-start.sh config/server.1.properties
bin/kafka-server-start.sh config/server.2.properties
bin/kafka-server-start.sh config/server.3.properties
```

### Create Topics 
```
bin/kafka-topics.sh --create --topic my-kafka-topic --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2
```

### Listing Kafka Topics 

```
bin/kafka-topics.sh --list --bootstrap-server localhost:9093
```

```
bin/kafka-topics.sh --describe --topic my-kafka-topic --bootstrap-server localhost:9093
```

Output
```
Topic: my-kafka-topic	TopicId: l3m1fgFbTL2JjTcsyoslRA	PartitionCount: 3	ReplicationFactor: 2	Configs: 
	Topic: my-kafka-topic	Partition: 0	Leader: 1	Replicas: 1,3	Isr: 1,3
	Topic: my-kafka-topic	Partition: 1	Leader: 2	Replicas: 2,1	Isr: 2,1
	Topic: my-kafka-topic	Partition: 2	Leader: 3	Replicas: 3,2	Isr: 3,2
```

## Producers: Sending Messages to a Topic

```
bin/kafka-console-producer.sh --broker-list localhost:9093,localhost:9094,localhost:9095 --topic my-kafka-topic
```

## Consumers
```
bin/kafka-console-consumer.sh --bootstrap-server localhost:9093 --topic my-kafka-topic --from-beginning
```

## Testing Replication: What if a Broker Goes Offline?
```
bin/kafka-console-consumer.sh --bootstrap-server localhost:9093 --topic my-kafka-topic --from-beginning --group group2
```




## References:

1. [Apache Kafka](https://kafka.apache.org/)
2. [Kafka Go Library](https://github.com/segmentio/kafka-go)
3. [How to Install and Run a Kafka Cluster Locally](https://www.sohamkamani.com/install-and-run-kafka-locally/)
4. [Kafka Tar File](https://archive.apache.org/dist/kafka/3.6.0/kafka_2.13-3.6.0.tgz)
5. [Setup Kakfa Locally](https://www.sohamkamani.com/install-and-run-kafka-locally/)
