package internal

type jsonTranslator struct{}

func (j *jsonTranslator) equals(other *translator) bool {
	this := translator(j)
	sameType := translatorEquals(&this, other)
	return sameType
}

func (j *jsonTranslator) encode(gum *gum) (raw, error) {
	panic("implement me")
}

func (j *jsonTranslator) decode(raw *raw) (gum, error) {
	panic("implement me")
}
