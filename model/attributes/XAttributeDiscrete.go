package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeDiscrete struct {
	XAttribute

	Value int

	Key        string
	Extensions []model.XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeDiscrete) XAttributeMarker() {}

func NewXAttributeDiscrete(key string, value int, extensions []model.XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
