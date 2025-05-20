package entities

type MessageEntity struct {
	ID      int
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

func NewMessageEntity(from string, to string, content string, kind ContentKind) *MessageEntity {
	return &MessageEntity{
		From:    from,
		To:      to,
		Content: content,
		Kind:    kind,
	}
}
