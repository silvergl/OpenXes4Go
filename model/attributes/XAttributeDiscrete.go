package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeDiscrete struct {
	XAttribute

	value int
}

func NewXAttributeDiscrete(key string, value int, extensions []model.XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		XAttribute: XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
