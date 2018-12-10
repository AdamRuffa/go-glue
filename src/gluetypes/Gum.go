package gluetypes

import (
	"go-glue/src/transl"
)

/*

A decoded representation of a glue call.

*/

type Gum interface {
	Translate(transl.Translator)
}
