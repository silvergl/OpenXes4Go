package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeContinuos struct {
	XAttribute
	Value float64

	Key        string
	Extensions []model.XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeContinuos) XAttributeMarker() {}

func NewXAttributeContinuos(key string, value float64, extensions []model.XExtension) (*XAttributeContinuos, error) {
	attr := &XAttributeContinuos{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
