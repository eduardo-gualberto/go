package main

import (
	"github.com/eduardo-gualberto/go.git/core/usecases"
	gatewaysdx "github.com/eduardo-gualberto/go.git/gateways/dx"
	infradx "github.com/eduardo-gualberto/go.git/infra/dx"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	fx.New(
		infradx.Module,
		gatewaysdx.Module,
		fx.Invoke(func(p usecases.CreateParticipant, u usecases.CreateUser) {
		}),
	).Run()
}
