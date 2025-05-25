package main

import (
	"fmt"

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
		fx.Invoke(func(p usecases.CreateParticipant, u usecases.ListUsers, l usecases.ListParticipants) {
			users, err := u.Execute()
			if err != nil {
				panic(err)
			}
			for _, u := range users {
				if u != nil {
					fmt.Printf("User %d: (%s, %s)\n", u.ID, u.Name, u.Number)
					continue
				}
				fmt.Println("Nil value passed")
			}
			eduardo := users[0]
			if err = p.Execute("5534991773442", "participante do eduardo", eduardo.ID); err != nil {
				panic(err)
			}
			participants, err := l.Execute()
			if err != nil {
				panic(err)
			}
			for _, p := range participants {
				if u != nil {
					fmt.Printf("Participant %d, of %d: (%s, %s)\n", p.ID, p.UserID, p.Name, p.Number)
					continue
				}
				fmt.Println("Nil value passed")
			}
		}),
	).Run()
}
