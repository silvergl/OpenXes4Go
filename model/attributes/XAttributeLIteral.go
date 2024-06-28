package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeLiteral struct {
	XAttribute

	Value string `xml:"value,attr"`

	Key        string
	Extensions []model.XExtension
	attributes map[string]XAttribute
}

func NewXAttributeLiteral(key string, value string, extensions []model.XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
