package attributes

import (
	"github.com/silvergl/OpenXes4Go/model/extensions"
)

type XAttributeDiscrete struct {
	XAttribute

	Value int

	Key        string
	Extensions []extensions.XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeDiscrete) XAttributeMarker() {}

func NewXAttributeDiscrete(key string, value int, extensions []extensions.XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
