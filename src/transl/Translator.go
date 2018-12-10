package transl

import (
	"go-glue/src/gluetypes"
	"go-glue/src/transl/impl"
)

/*

An interface for all go-glue translators.

*/

type Translator interface {
	Encode(gluetypes.Gum) (gluetypes.Raw, error)
	Decode(gluetypes.Raw) (gluetypes.Gum, error)
}

const JsonEncoding = "application/json"
const XmlEncoding = "application/xml"
const ProtobufEncoding = "application/protobuf"
const UndefinedEncoding = "undefined"

func NewTranslator(encoding string) Translator {
	switch encoding {
	case JsonEncoding:
		return &impl.JsonTranslator{}
	case ProtobufEncoding:
		return &impl.ProtobufTranslator{}
	case XmlEncoding:
		fallthrough
	case "text/xml":
		return &impl.XmlTranslator{}
	case UndefinedEncoding:
		fallthrough
	default:
		return &impl.UndefinedTranslator{}
	}
}
