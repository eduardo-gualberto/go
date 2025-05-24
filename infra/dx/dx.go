package dx

import (
	"github.com/eduardo-gualberto/go.git/infra/db"
	"github.com/eduardo-gualberto/go.git/infra/httpclient"
	"github.com/eduardo-gualberto/go.git/infra/wabaapi"
	"go.uber.org/fx"
)

var Module = fx.Module("infra",
	fx.Provide(fx.Annotate(
		httpclient.GetWabaHttpClient,
		fx.ResultTags(`name:"http_waba"`),
	)),
	fx.Provide(fx.Annotate(
		wabaapi.NewWabaApi,
		fx.ParamTags(`name:"http_waba"`),
	)),
	fx.Provide(db.NewAppDbConn),
)
