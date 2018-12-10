package src

type JsonTranslator struct{}

func (JsonTranslator) Encode(Gum) (Raw, error) {
	panic("implement me")
}

func (JsonTranslator) Decode(Raw) (Gum, error) {
	panic("implement me")
}
