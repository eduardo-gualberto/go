package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var once = &sync.Once{}
var appDbInstance *pgx.Conn

func getDbUrl() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	pwd := os.Getenv("DB_PASSWORD")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pwd, host, port, name)
	return url
}

func NewAppDbConn() *pgx.Conn {
	once.Do(func() {
		url := getDbUrl()
		c, err := pgx.Connect(context.Background(), url)
		if err != nil {
			panic(err)
		}
		appDbInstance = c
	})
	return appDbInstance
}
