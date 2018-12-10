package impl

import (
	"go-glue/src/gluetypes"
)

type XmlTranslator struct{}

func (XmlTranslator) Encode(gluetypes.Gum) (gluetypes.Raw, error) {
	panic("implement me")
}

func (XmlTranslator) Decode(gluetypes.Raw) (gluetypes.Gum, error) {
	panic("implement me")
}
