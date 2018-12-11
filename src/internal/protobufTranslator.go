package internal

type protobufTranslator struct{}

func (protobufTranslator) equals(other *translator) bool {
	panic("implement me")
}

func (protobufTranslator) encode(gum *gum) (raw, error) {
	panic("implement me")
}

func (protobufTranslator) decode(raw *raw) (gum, error) {
	panic("implement me")
}
