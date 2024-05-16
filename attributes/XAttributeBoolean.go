package attributes

import (
	"github.com/openxes4go/model"
)

type XAttributeBoolean struct {
	model.XAttribute
	value bool
}

/*
 * New XAttributeBoolean
 */
func NewXAttributeBoolean(key string, value bool, extensions []model.XExtension) (*XAttributeBoolean, error) {
	attr := &XAttributeBoolean{
		XAttribute: model.XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}
	return attr, nil

}
