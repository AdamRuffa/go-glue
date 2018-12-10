package src

/*

An interface for all go-glue translators.

*/

type Translator interface {
	Encode(Gum) (Raw, error)
	Decode(Raw) (Gum, error)
}

const JsonEncoding = "application/json"
const XmlEncoding = "application/xml"
const ProtobufEncoding = "application/protobuf"
const UndefinedEncoding = "undefined"

func NewTranslator(encoding string) Translator {
	switch encoding {
	case JsonEncoding:
		return &JsonTranslator{}
	case ProtobufEncoding:
		return &ProtobufTranslator{}
	case XmlEncoding:
		fallthrough
	case "text/xml":
		return &XmlTranslator{}
	case UndefinedEncoding:
		fallthrough
	default:
		return &UndefinedTranslator{}
	}
}
