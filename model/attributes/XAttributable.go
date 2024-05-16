package attributes

type XAttributable struct {
	attributes []XAttribute
}

func NewXAtributable(attributes []XAttribute) (*XAttributable, error) {
	attr := &XAttributable{
		attributes: attributes,
	}
	return attr, nil
}
