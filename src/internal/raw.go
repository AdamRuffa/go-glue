package internal

/*

An interface and struct for raw data representation.

*/

type raw interface {
	Translate() (gum, error)
}

func NewRaw() raw {
	return &rawImpl{}
}

type rawImpl struct {
	bytes          *[]byte
	text           *string
	encodingScheme *string
	translator     *translator
}

func (raw *rawImpl) getTranslator() translator           { return *raw.translator }
func (raw *rawImpl) setTranslator(translator translator) { raw.translator = &translator }
func (raw *rawImpl) resolveTranslator() {
	translator := raw.getTranslator()

	if translator == nil {

	}
}
func (raw *rawImpl) Translate() (gum gum, e error) {
	panic("implement me")
}
