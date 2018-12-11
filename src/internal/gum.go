package internal

import (
	"errors"
	"reflect"
)

/*

A decoded representation of a glue call.

*/

type gum interface {
	unmarshal(*interface{}) error
	getTranslator() translator
	setTranslator(translator)
	translate(...translator) (*raw, error)
}

type gumField struct {
	field     string
	class     reflect.Type
	subFields map[string]*gumField
	parent    *gumField
}

type gumImpl struct {
	root       *gumField
	translator *translator
}

const ROOT string = "/"

func newGum() gum {
	root := &gumField{
		new(string(ROOT)),
		nil,
		make(map[string]*gumField),
		nil,
	}
	root.parent = root
	defaultTranslator := newTranslator(UndefinedEncoding)
	return &gumImpl{
		root,
		&defaultTranslator,
	}
}

func (g gumImpl) getTranslator() translator           { return *g.translator }
func (g gumImpl) setTranslator(translator translator) { g.translator = &translator }
func (g gumImpl) unmarshal(object *interface{}) (e error) {
	gum := gum(g)
	if object == nil {
		e = errors.New("gum.unmarshal expects a nonnull object reference")
		return
	}
	panic(gum)

}
func (g gumImpl) translate(...translator) (rawRef *raw, e error) {
	gum := gum(g)
	translator := g.getTranslator()
	if translator == nil {
		e = errors.New("gum.unmarshal expects translator to first be set")
		return
	}
	raw, e := translator.encode(&gum)
	return &raw, e
}
