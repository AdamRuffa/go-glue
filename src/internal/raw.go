package internal

/*

An interface and struct for raw data representation.

*/

type raw interface {
	GetTranslator() translator
	SetTranslator(translator translator)
	Translate() (gum, error)
}

func NewRaw() raw {
	return &rawImpl{}
}

type rawImpl struct{}

func (raw *rawImpl) GetTranslator() translator {
	panic("implement me")
}

func (raw *rawImpl) SetTranslator(translator translator) {
	panic("implement me")
}

func (raw *rawImpl) Translate() (gum, error) {
	panic("implement me")
}
