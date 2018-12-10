package impl

import (
	"go-glue/src/gluetypes"
)

type UndefinedTranslator struct{}

func (UndefinedTranslator) Encode(gluetypes.Gum) (gluetypes.Raw, error) {
	panic("implement me")
}

func (UndefinedTranslator) Decode(gluetypes.Raw) (gluetypes.Gum, error) {
	panic("implement me")
}
