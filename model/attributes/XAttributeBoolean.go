package attributes

import (
	"github.com/silvergl/OpenXes4Go/model"
)

type XAttributeBoolean struct {
	XAttribute
	value bool
}

/*
 * New XAttributeBoolean
 */
func NewXAttributeBoolean(key string, value bool, extensions []model.XExtension) (*XAttributeBoolean, error) {
	attr := &XAttributeBoolean{
		XAttribute: XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}
	return attr, nil

}
