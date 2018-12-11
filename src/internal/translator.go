package internal

/*

An interface for all go-glue translators.

*/

type translator interface {
	encode(gum *gum) (raw, error)
	decode(raw *raw) (gum, error)
}

const JsonEncoding = "application/json"
const XmlEncoding = "application/xml"
const ProtobufEncoding = "application/protobuf"
const UndefinedEncoding = "undefined"

func newTranslator(encoding string) translator {
	switch encoding {
	case JsonEncoding:
		return &jsonTranslator{}
	case ProtobufEncoding:
		return &protobufTranslator{}
	case XmlEncoding:
		fallthrough
	case "text/xml":
		return &xmlTranslator{}
	case UndefinedEncoding:
		fallthrough
	default:
		return &undefinedTranslator{}
	}
}
