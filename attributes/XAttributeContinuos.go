package attributes

import "github.com/openxes4go/model"

type XAttributeContinuos struct {
	model.XAttribute
	value float64
}

func NewXAttributeContinuos(key string, value float64, extensions []model.XExtension) (*XAttributeContinuos, error) {
	attr := &XAttributeContinuos{
		XAttribute: model.XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
