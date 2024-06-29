package attributes

import (
	"github.com/silvergl/OpenXes4Go/model/extensions"
)

type XAttributeTimeStamp struct {
	XAttribute

	value string

	Key        string
	Extensions []extensions.XExtension
	attributes map[string]XAttribute
}

func NewXAttributeDTimeStamp(key string, value string, extensions []extensions.XExtension) (*XAttributeTimeStamp, error) {

	attr := &XAttributeTimeStamp{
		Key:        key,
		Extensions: extensions,
		value:      value,
	}

	return attr, nil
}
