package attributes

import "github.com/openxes4go/model"

type XAttributeLiteral struct {
	model.XAttribute

	value string
}

func NewXAttributeLiteral(key string, value string, extensions []model.XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		XAttribute: model.XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
