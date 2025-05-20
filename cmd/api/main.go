package main

import (
	"fmt"
	"net/http"

	gatewaysdx "github.com/eduardo-gualberto/go.git/gateways/dx"
	"github.com/eduardo-gualberto/go.git/gateways/handlers/wabahandler"
	infradx "github.com/eduardo-gualberto/go.git/infra/dx"
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
		gatewaysdx.Module,
		fx.Invoke(func(wh *wabahandler.WabaHandler) {
			http.HandleFunc("GET /webhook", wh.HandleAuthenticate)
			http.HandleFunc("POST /webhook", wh.HandleMessage)
			http.ListenAndServe(":8080", nil)
		}),
	).Run()
}
