package src

/*

An interface and struct for raw data representation.

*/

type Raw interface {
	GetTranslator() Translator
	SetTranslator(translator Translator)
	Translate() (Gum, error)
}

func NewRaw() Raw {
	return &rawImpl{}
}

type rawImpl struct{}

func (rawImpl) GetTranslator() Translator {
	panic("implement me")
}

func (rawImpl) SetTranslator(translator Translator) {
	panic("implement me")
}

func (rawImpl) Translate() (Gum, error) {
	panic("implement me")
}
