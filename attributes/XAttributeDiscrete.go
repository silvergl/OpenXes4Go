package attributes

import "github.com/openxes4go/model"

type XAttributeDiscrete struct {
	model.XAttribute

	value int
}

func NewXAttributeDiscrete(key string, value int, extensions []model.XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		XAttribute: model.XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
