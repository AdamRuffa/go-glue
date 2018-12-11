package internal

type undefinedTranslator struct{}

func (undefinedTranslator) Encode(gum) (raw, error) {
	panic("implement me")
}

func (undefinedTranslator) Decode(raw) (gum, error) {
	panic("implement me")
}
