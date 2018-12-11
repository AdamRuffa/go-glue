package internal

/*

An interface and struct for raw data representation.

*/

type raw interface {
	GetTranslator() translator
	SetTranslator(translator translator)
	Translate() (gum, error)
}

func newRaw() raw {
	return &rawImpl{}
}

type rawImpl struct{}

func (rawImpl) GetTranslator() translator {
	panic("implement me")
}

func (rawImpl) SetTranslator(translator translator) {
	panic("implement me")
}

func (rawImpl) Translate() (gum, error) {
	panic("implement me")
}
