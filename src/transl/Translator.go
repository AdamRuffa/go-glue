package transl

import (
	"go-glue/src/types"
)

/*

An interface for all go-glue translators.

*/

type Translator interface {
	encode(types.Gum) (types.Raw, error)
	decode(types.Raw) (types.Gum, error)
}
