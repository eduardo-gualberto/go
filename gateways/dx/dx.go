package dx

import (
	"github.com/eduardo-gualberto/go.git/core/interfaces"
	"github.com/eduardo-gualberto/go.git/core/usecases"
	"github.com/eduardo-gualberto/go.git/gateways/handlers/wabahandler"
	"github.com/eduardo-gualberto/go.git/gateways/interfaceimpls/chatcompletionllms"
	"github.com/eduardo-gualberto/go.git/gateways/interfaceimpls/messagesenders"
	"github.com/eduardo-gualberto/go.git/gateways/usecaseimpls"
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
	fx.Provide(fx.Annotate(
		usecaseimpls.NewRespondToUserImpl,
		fx.ParamTags(`name:"chatcompletion_openai"`, `name:"sender_openai"`),
		fx.As(new(usecases.RespondToUser)),
	)),
)
