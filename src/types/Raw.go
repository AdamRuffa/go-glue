package types

import (
	"go-glue/src/transl"
)

/*

An interface and struct for raw data representation.

*/

type Raw interface {
	Translate(transl.Translator)
}
