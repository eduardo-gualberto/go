package dx

import (
	"github.com/eduardo-gualberto/go.git/core/interfaces"
	"github.com/eduardo-gualberto/go.git/gateways/chatcompletionllms"
	"github.com/eduardo-gualberto/go.git/gateways/handlers/wabahandler"
	"github.com/eduardo-gualberto/go.git/gateways/messagesenders"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"gateways",
	fx.Provide(wabahandler.NewWabaHandler),
	fx.Provide(fx.Annotate(
		chatcompletionllms.NewOpenAiLLM,
		fx.As(new(interfaces.ChatCompletionLLM)),
		fx.ResultTags(`name:"chatcompletion_openai"`),
	)),
	fx.Provide(fx.Annotate(
		messagesenders.NewWabaMessageSender,
		fx.As(new(interfaces.MessageSender)),
		fx.ResultTags(`name:"sender_openai"`),
	)),
)
