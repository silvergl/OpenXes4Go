package model

type XAttributable struct {
	attributes []XAttribute
}

func NewXAttributable(attributes []XAttribute) (*XAttributable, error) {
	attr := &XAttributable{
		attributes: attributes,
	}
	return attr, nil
}
