package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeTimeStamp struct {
	XAttribute

	value string

	Key        string
	Extensions []model.XExtension
	attributes map[string]XAttribute
}

func NewXAttributeDTimeStamp(key string, value string, extensions []model.XExtension) (*XAttributeTimeStamp, error) {

	attr := &XAttributeTimeStamp{
		Key:        key,
		Extensions: extensions,
		value:      value,
	}

	return attr, nil
}
