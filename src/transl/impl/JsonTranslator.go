package impl

import (
	"go-glue/src/gluetypes"
)

type JsonTranslator struct{}

func (JsonTranslator) Encode(gluetypes.Gum) (gluetypes.Raw, error) {
	panic("implement me")
}

func (JsonTranslator) Decode(gluetypes.Raw) (gluetypes.Gum, error) {
	panic("implement me")
}
