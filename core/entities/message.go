package entities

type MessageEntity struct {
	From    string
	To      string
	Content string
	Kind    ContentKind
}

type ContentKind string

var (
	TextContentKind     ContentKind = "text"
	AudioContentKind    ContentKind = "audio"
	ImageContentKind    ContentKind = "image"
	DocumentContentKind ContentKind = "document"
)
