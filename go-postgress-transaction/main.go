package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Create a new connection to our database
	connStr := "user=postgres dbname=postgres password=admin123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return
	}

	// Run a query to get a count of all cats
	row := tx.QueryRow("SELECT count(*) FROM pets WHERE species='cat'")
	var catCount int
	// Store the count in the `catCount` variable
	err = row.Scan(&catCount)
	if err != nil {
		tx.Rollback()
		return
	}

	// Now update the food table, increasing the quantity of cat food by 10x the number of cats
	_, err = tx.ExecContext(ctx, "UPDATE food SET quantity=quantity+$1 WHERE name='Cat Food'", 10*catCount)
	if err != nil {
		tx.Rollback()
		return
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
