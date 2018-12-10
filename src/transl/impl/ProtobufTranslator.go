package impl

import (
	"go-glue/src/gluetypes"
)

type ProtobufTranslator struct{}

func (ProtobufTranslator) Encode(gluetypes.Gum) (gluetypes.Raw, error) {
	panic("implement me")
}

func (ProtobufTranslator) Decode(gluetypes.Raw) (gluetypes.Gum, error) {
	panic("implement me")
}
