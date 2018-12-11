package internal

type protobufTranslator struct{}

func (protobufTranslator) Encode(gum) (raw, error) {
	panic("implement me")
}

func (protobufTranslator) Decode(raw) (gum, error) {
	panic("implement me")
}
