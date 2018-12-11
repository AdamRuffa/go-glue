package internal

import (
	"errors"
	"reflect"
)

/*

A decoded representation of a glue call.

*/

type gum interface {
	Unmarshal(*interface{}) error
	GetTranslator() translator
	SetTranslator(translator)
	Translate(...translator) (*raw, error)
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

func root() string { return "/" }

func NewGum() gum {
	root := &gumField{
		root(),
		nil,
		make(map[string]*gumField),
		nil,
	}
	root.parent = root
	defaultTranslator := NewTranslator(UndefinedEncoding)
	return &gumImpl{
		root,
		&defaultTranslator,
	}
}

func (g *gumImpl) GetTranslator() translator           { return *g.translator }
func (g *gumImpl) SetTranslator(translator translator) { g.translator = &translator }
func (g *gumImpl) Unmarshal(object *interface{}) (e error) {
	gum := gum(g)
	if object == nil {
		e = errors.New("gum.unmarshal expects a nonnull object reference")
		return
	}
	panic(gum)

}
func (g *gumImpl) Translate(...translator) (rawRef *raw, e error) {
	gum := gum(g)
	translator := g.GetTranslator()
	if translator == nil {
		e = errors.New("gum.unmarshal expects translator to first be set")
		return
	}
	raw, e := translator.encode(&gum)
	return &raw, e
}
