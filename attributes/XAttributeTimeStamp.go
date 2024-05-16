package attributes

import "github.com/openxes4go/model"

type XAttributeTimeStamp struct {
	model.XAttribute

	value string
}

func NewXAttributeDTimeStamp(key string, value string, extensions []model.XExtension) (*XAttributeTimeStamp, error) {
	attr := &XAttributeTimeStamp{
		XAttribute: model.XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
