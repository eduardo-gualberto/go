package wabaapi

import "os"

type textOpts struct {
	body string
	from string
}
type TextOpt func(*textOpts)

func (wpp *WabaApi) WithBody(body string) TextOpt {
	return func(o *textOpts) { o.body = body }
}

func (w *WabaApi) SendText(to string, opts ...TextOpt) (bool, error) {
	o := textOpts{
		from: os.Getenv("META_PHONE_NUMBER_ID"),
		body: "",
	}

	for _, fn := range opts {
		fn(&o)
	}
	return w.sendMessage(o.from, to, "text", map[string]any{
		"body": o.body,
	})
}
