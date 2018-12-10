package src

type ProtobufTranslator struct{}

func (ProtobufTranslator) Encode(Gum) (Raw, error) {
	panic("implement me")
}

func (ProtobufTranslator) Decode(Raw) (Gum, error) {
	panic("implement me")
}
