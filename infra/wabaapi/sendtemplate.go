package wabaapi

import "os"

type templateOpts struct {
	name     string
	language string
	from     string
}
type TemplateOpt func(*templateOpts)

func (wpp *WabaApi) WithLanguage(code string) TemplateOpt {
	return func(o *templateOpts) { o.language = code }
}
func (wpp *WabaApi) WithName(name string) TemplateOpt {
	return func(o *templateOpts) { o.name = name }
}

func (w *WabaApi) SendTemplate(to string, opts ...TemplateOpt) (bool, error) {
	// defaults
	o := templateOpts{
		from:     os.Getenv("META_PHONE_NUMBER_ID"),
		language: "en_US",
		name:     "hello_world",
	}
	// apply passed opts
	for _, fn := range opts {
		fn(&o)
	}

	return w.sendMessage(o.from, to, "template", map[string]any{
		"name": o.name,
		"language": map[string]any{
			"code": o.language,
		},
	})
}
