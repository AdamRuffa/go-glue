package internal

type jsonTranslator struct{}

func (jsonTranslator) encode(gum *gum) (raw, error) {
	panic("implement me")
}

func (jsonTranslator) decode(raw *raw) (gum, error) {
	panic("implement me")
}
