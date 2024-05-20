package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Bird struct {
	id          int
	Species     string
	Description string
}

// Ref : https://www.sohamkamani.com/golang/sql-database/
func main() {
	fmt.Println("MySQl Database example in golang ")
	db, err := sql.Open("mysql", "root:admin123@/bird_encyclopedia")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// db.SetConnMaxLifetime() is required to ensure connections are closed by the driver safely before connection is closed by MySQL server, OS, or other middlewares.
	// Since some middlewares close idle connections by 5 minutes, we recommend timeout shorter than 5 minutes.
	// This setting helps load balancing and changing system variables too.
	db.SetConnMaxLifetime(3 * time.Minute)

	// db.SetMaxOpenConns() is highly recommended to limit the number of connection used by the application.
	// There is no recommended limit number because it depends on application and MySQL server.
	db.SetMaxOpenConns(10)

	// db.SetMaxIdleConns() is recommended to be set same to db.SetMaxOpenConns().
	// When it is smaller than SetMaxOpenConns(), connections can be opened and closed much more frequently than you expect.
	// Idle connections can be closed by the db.SetConnMaxLifetime().
	// If you want to close idle connections more rapidly, you can use db.
	// SetConnMaxIdleTime() since Go 1.15.
	db.SetMaxIdleConns(10)

	// Idle Connection Timeout
	db.SetConnMaxIdleTime(10 * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	// Single row Databse query with limit 1
	r := db.QueryRow("SELECT bird, description FROM birds LIMIT 1")
	bird := Bird{}

	if err := r.Scan(&bird.Species, &bird.Description); err != nil {
		log.Fatalf("could not scan row: %v", err)
	}

	fmt.Printf("found bird: %+v\n", bird)

	// Select Query
	rows, err := db.Query("SELECT * FROM birds;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	birds := []Bird{}
	for rows.Next() {
		bird := Bird{}
		if err := rows.Scan(&bird.id, &bird.Species, &bird.Description); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		log.Println(bird)
		birds = append(birds, bird)
	}
	fmt.Printf("found %d birds: %+v", len(birds), birds)

	// Select Query with Where Caluse
	birdName := "eagle"
	fmt.Sprintf("SELECT bird, description FROM birds WHERE bird=\"%s\" LIMIT %d", birdName, 1)
	row := db.QueryRow(fmt.Sprintf("SELECT bird, description FROM birds WHERE bird=\"%s\" LIMIT %d", birdName, 1))

	if err := row.Scan(&bird.Species, &bird.Description); err != nil {
		log.Fatalf("could not scan row: %v", err)
	}
	log.Println(bird)

	// Executing Writes - INSERT, UPDATE, and DELETE

	// sample data that we want to insert
	newBird := Bird{
		Species:     "Peigon",
		Description: "Weather Air is good",
	}

	// the `Exec` method returns a `Result` type instead of a `Row`
	// we follow the same argument pattern to add query params

	stmt, err := db.Prepare("INSERT INTO birds (bird, description) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(newBird.Species, newBird.Description)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!", id)

	// // we can log how many rows were inserted
	// fmt.Println("inserted", rowsAffected, "rows")

	// Connection Pooling - Timeouts and Max/Idle Connections

	// Maximum Idle Connections
	// db.SetMaxIdleConns(5)
	// Maximum Open Connections
	// db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	// db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	// db.SetConnMaxLifetime(30 * time.Second)

	// Query Timeouts - Using Context Cancellation
	// create a parent context
	ctx := context.Background()
	// create a context from the parent context with a 300ms timeout
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)
	// The context variable is passed to the `QueryContext` method as
	// the first argument
	// the pg_sleep method is a function in Postgres that will halt for
	// the provided number of seconds. We can use this to simulate a
	// slow query
	_, err = db.QueryContext(ctx, "SELECT * from birds")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
}
