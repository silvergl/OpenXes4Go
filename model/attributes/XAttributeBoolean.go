package attributes

import (
	"github.com/silvergl/OpenXes4Go/model"
)

type XAttributeBoolean struct {
	XAttribute
	Value      bool
	Key        string
	Extensions []model.XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeBoolean) XAttributeMarker() {}

/*
 * New XAttributeBoolean
 */
func NewXAttributeBoolean(key string, value bool, extensions []model.XExtension) (*XAttributeBoolean, error) {
	attr := &XAttributeBoolean{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}
	return attr, nil

}
