package internal

import "reflect"

/*

A decoded representation of a glue call.

*/

type gum interface {
	Unmarshal(*interface{}) error
	GetTranslator() translator
	SetTranslator(translator)
	Translate(translator)
}

type gumField struct {
	field     string
	class     reflect.Type
	subFields map[string]*gumField
	parent    *gumField
}

type gumImpl struct {
	root *gumField
}

const ROOT string = "/"

func newGum() gum {
	root := &gumField{
		string(ROOT),
		nil,
		make(map[string]*gumField),
		nil,
	}
	root.parent = root
	return &gumImpl{
		root,
	}
}

func (gumImpl) Unmarshal(*interface{}) error {
	panic("implement me")
}

func (gumImpl) GetTranslator() translator {
	panic("implement me")
}

func (gumImpl) SetTranslator(translator translator) {
	panic("implement me")
}

func (gumImpl) Translate(translator) {
	panic("implement me")
}
