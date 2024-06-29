package attributes

type XAttributeLiteral struct {
	XAttribute

	Value string `xml:"value,attr"`

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func NewXAttributeLiteral(key string, value string, extensions []XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
