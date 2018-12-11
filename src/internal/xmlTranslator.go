package internal

type xmlTranslator struct{}

func (xmlTranslator) Encode(gum) (raw, error) {
	panic("implement me")
}

func (xmlTranslator) Decode(raw) (gum, error) {
	panic("implement me")
}
