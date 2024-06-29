package attributes

type XAttributable struct {
	Attributes []XAttribute
}git

func NewXAttributable(attributes []XAttribute) (*XAttributable, error) {
	attr := &XAttributable{
		Attributes: attributes,
	}
	return attr, nil
}
