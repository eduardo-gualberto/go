package main

import (
	"context"
	"fmt"
	"os"

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

	var result string
	if err = conn.QueryRow(context.Background(), "select version();").Scan(&result); err != nil {
		fmt.Printf("Error querying pg: %v", err)
		return
	}
	fmt.Printf("Result is: %s\n", result)
}
