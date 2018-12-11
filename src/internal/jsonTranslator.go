package internal

type jsonTranslator struct{}

func (jsonTranslator) Encode(gum) (raw, error) {
	panic("implement me")
}

func (jsonTranslator) Decode(raw) (gum, error) {
	panic("implement me")
}
