package attributes

import "github.com/openxes4go/model"

type XAttributeID struct {
	model.XAttribute

	value string //TODO ID OWN TYPE
}

func NewXAttributeID(key string, value string, extensions []model.XExtension) (*XAttributeID, error) {
	attr := &XAttributeID{
		XAttribute: model.XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
