package internal

import (
	"errors"
	"fmt"
)

type undefinedTranslator struct{}

func (undefinedTranslator) equals(other *translator) bool {
	panic("implement me")
}

func (undefinedTranslator) encode(gum *gum) (raw, error) {
	return nil, errors.New(fmt.Sprintf(
		"translator for gum %v is undefined; "+
			"please set a known encodingSchme", gum))
}

func (undefinedTranslator) decode(raw *raw) (gum, error) {
	return nil, errors.New(fmt.Sprintf(
		"translator for gum %v is undefined; "+
			"please set a known encodingSchme", raw))
}
