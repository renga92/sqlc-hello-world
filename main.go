package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/renga92/sqlc-hello-world/db" // Select the folder to get the DB models
	"log"
	"os"
	// "time"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("DB Connection successful!")
	}

	queries := db.New(conn)
	fmt.Println(queries)

	insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(insertedAuthor)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(authors)
}
