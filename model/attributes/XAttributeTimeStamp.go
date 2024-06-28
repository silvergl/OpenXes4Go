package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeTimeStamp struct {
	XAttribute

	value string
}

func NewXAttributeDTimeStamp(key string, value string, extensions []model.XExtension) (*XAttributeTimeStamp, error) {
	attr := &XAttributeTimeStamp{
		XAttribute: XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
