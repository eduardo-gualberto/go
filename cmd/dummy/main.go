package main

import (
	"fmt"

	appdx "github.com/eduardo-gualberto/go.git/internal/application/dx"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
	infradx "github.com/eduardo-gualberto/go.git/internal/infra/dx"
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
		appdx.Module,
		fx.Invoke(func(ce usecases.CreateEvent, le usecases.ListEvents, co usecases.CreateOccurrence, lo usecases.ListOccurrences) {
			fmt.Println("Hello, world!")
		}),
	).Run()
}
