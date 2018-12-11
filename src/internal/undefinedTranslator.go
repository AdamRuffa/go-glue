package internal

type undefinedTranslator struct{}

func (undefinedTranslator) encode(gum *gum) (raw, error) {
	panic("implement me")
}

func (undefinedTranslator) decode(raw *raw) (gum, error) {
	panic("implement me")
}
