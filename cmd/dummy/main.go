package main

import (
	"context"
	"fmt"
	"os"

	db "github.com/eduardo-gualberto/go.git/infra/db/gen"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	pwd := os.Getenv("DB_PASSWORD")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pwd, host, port, name)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Error connecting to pg: %v", err)
		return
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	version, err := q.GetVersion(context.Background())
	if err != nil {
		fmt.Printf("error getting version: %v", err)
		return
	}
	fmt.Print(version)
}
