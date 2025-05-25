package main

import (
	"fmt"
	"net/http"

	appdx "github.com/eduardo-gualberto/go.git/internal/application/dx"
	"github.com/eduardo-gualberto/go.git/internal/application/handlers/wabahandler"
	infradx "github.com/eduardo-gualberto/go.git/internal/infra/dx"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading env: %v", err)
	}
}

func main() {
	fx.New(
		infradx.Module,
		appdx.Module,
		fx.Invoke(func(wh *wabahandler.WabaHandler) {
			http.HandleFunc("GET /webhook", wh.HandleAuthenticate)
			http.HandleFunc("POST /webhook", wh.HandleMessage)
			http.ListenAndServe(":8080", nil)
		}),
	).Run()
}
