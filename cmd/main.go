package main

import (
	"context"
	"log"
	"net/http"

	"github.com/akubi0w1/ent-sample/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbClient, err := ent.Open("mysql", "worker:password@tcp(db:3306)/main")
	if err != nil {
		log.Fatalf("failed to open connetion to mysql: %v", err)
	}
	defer dbClient.Close()

	// migration
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema resources: %v", err)
	}

	// create user
	// now := time.Now()
	// u, err := dbClient.User.
	// 	Create().
	// 	SetCreatedAt(now).
	// 	SetUpdatedAt(now).
	// 	SetAccountID("accountID").
	// 	Save(context.Background())
	// if err != nil {
	// 	log.Printf("failed to create user: %v", err)
	// }
	// log.Println("user created: ", u)

	http.ListenAndServe(":8080", nil)
}
