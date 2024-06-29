package attributes

import (
	"github.com/silvergl/OpenXes4Go/model/extensions"
)

type XAttributeID struct {
	XAttribute

	Value string //TODO ID OWN TYPE

	Key        string
	Extensions []extensions.XExtension
	attributes map[string]XAttribute
}

func NewXAttributeID(key string, value string, extensions []extensions.XExtension) (*XAttributeID, error) {
	attr := &XAttributeID{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
