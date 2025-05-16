package dx

import (
	"github.com/eduardo-gualberto/go.git/core/usecases"
	"go.uber.org/fx"
)

var Module = fx.Module("core", fx.Provide(fx.Annotate(
	usecases.NewRespondToUser,
	fx.ParamTags(`name:"chatcompletion_openai"`, `name:"sender_openai"`),
)))
