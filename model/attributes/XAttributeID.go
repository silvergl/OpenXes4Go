package attributes

import "github.com/silvergl/OpenXes4Go/model"

type XAttributeID struct {
	XAttribute

	value string //TODO ID OWN TYPE
}

func NewXAttributeID(key string, value string, extensions []model.XExtension) (*XAttributeID, error) {
	attr := &XAttributeID{
		XAttribute: XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
