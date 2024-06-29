package attributes

import (
	"github.com/silvergl/OpenXes4Go/model/extensions"
)

type XAttributeLiteral struct {
	XAttribute

	Value string `xml:"value,attr"`

	Key        string
	Extensions []extensions.XExtension
	attributes map[string]XAttribute
}

func NewXAttributeLiteral(key string, value string, extensions []extensions.XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
