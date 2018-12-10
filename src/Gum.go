package src

/*

A decoded representation of a glue call.

*/

type Gum interface {
	GetTranslator() Translator
	SetTranslator(translator Translator)
	Translate(Translator)
}

func NewGum() Gum {
	return &gumImpl{}
}

type gumImpl struct{}

func (gumImpl) GetTranslator() Translator {
	panic("implement me")
}

func (gumImpl) SetTranslator(translator Translator) {
	panic("implement me")
}

func (gumImpl) Translate(Translator) {
	panic("implement me")
}
