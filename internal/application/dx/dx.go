package dx

import (
	"github.com/eduardo-gualberto/go.git/internal/application/gateways/interfaceimpls/chatcompletionllms"
	"github.com/eduardo-gualberto/go.git/internal/application/gateways/interfaceimpls/messagesenders"
	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/handlers/wabahandler"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/participantrepo"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/userrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"gateways",
	fx.Provide(wabahandler.NewWabaHandler),
	fx.Provide(fx.Annotate(
		chatcompletionllms.NewOpenAiLLM,
		fx.As(new(interfaces.ChatCompletionLLM)),
	)),
	fx.Provide(fx.Annotate(
		messagesenders.NewWabaMessageSender,
		fx.As(new(interfaces.MessageSender)),
	)),
	fx.Provide(fx.Annotate(
		usecaseimpls.NewRespondToUserImpl,
		fx.As(new(usecases.RespondToUser)),
	)),
	fx.Provide(fx.Annotate(
		participantrepo.NewParticipantRepoImpl,
		fx.As(new(participantrepo.ParticipantRepo)),
	)),
	fx.Provide(fx.Annotate(
		usecaseimpls.NewCreateParticipantImpl,
		fx.As(new(usecases.CreateParticipant)),
	)),
	fx.Provide(fx.Annotate(
		userrepo.NewUserRepoImpl,
		fx.As(new(userrepo.UserRepo)),
	)),
	fx.Provide(fx.Annotate(
		usecaseimpls.NewCreateUserImpl,
		fx.As(new(usecases.CreateUser)),
	)),
	fx.Provide(fx.Annotate(
		usecaseimpls.NewListUsersImpl,
		fx.As(new(usecases.ListUsers)),
	)),
	fx.Provide(fx.Annotate(
		usecaseimpls.NewListParticipantsImpl,
		fx.As(new(usecases.ListParticipants)),
	)),
)
