package src

type UndefinedTranslator struct{}

func (UndefinedTranslator) Encode(Gum) (Raw, error) {
	panic("implement me")
}

func (UndefinedTranslator) Decode(Raw) (Gum, error) {
	panic("implement me")
}
