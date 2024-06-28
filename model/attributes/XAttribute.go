package attributes

import (
	"github.com/silvergl/OpenXes4Go/model"
)

type XAttribute struct {
	XAttributable

	Key        string
	Extensions []model.XExtension
	attributes map[string]XAttribute
}

/*
 * Empty Attribute
 */
func NewXAttributeEmpty(key string) (*XAttribute, error) {
	attr := &XAttribute{
		Key: key,

		XAttributable: XAttributable{
			Attributes: make([]XAttribute, 0),
		},
	}
	return attr, nil
}

/*
 * Creates New Attribute
 */
func NewXAttribute(key string, extensions []model.XExtension) (*XAttribute, error) {
	attr := &XAttribute{
		Key:        "0",
		Extensions: extensions,
		XAttributable: XAttributable{
			Attributes: make([]XAttribute, 0),
		},
	}
	return attr, nil
}
