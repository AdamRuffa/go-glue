package internal

type xmlTranslator struct{}

func (xmlTranslator) encode(gum *gum) (raw, error) {
	panic("implement me")
}

func (xmlTranslator) decode(raw *raw) (gum, error) {
	panic("implement me")
}
