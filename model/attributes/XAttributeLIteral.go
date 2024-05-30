package attributes

import "github.com/openxes4go/model"

type XAttributeLiteral struct {
	XAttribute

	Value string `xml:"value,attr"`
}

func NewXAttributeLiteral(key string, value string, extensions []model.XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		XAttribute: XAttribute{
			Key:        key,
			Extensions: extensions,
		},
		Value: value,
	}

	return attr, nil
}
