package src

type XmlTranslator struct{}

func (XmlTranslator) Encode(Gum) (Raw, error) {
	panic("implement me")
}

func (XmlTranslator) Decode(Raw) (Gum, error) {
	panic("implement me")
}
