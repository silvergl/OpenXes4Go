package attributes

type XAttributeLiteral struct {
	XAttribute

	Value string `xml:"value,attr"`

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeLiteral) XAttributeMarker() {
}

func (attr XAttributeLiteral) GetExtensions() []XExtension {
	return attr.Extensions
}

func (attr XAttributeLiteral) GetAttributes() map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeLiteral) GetValue() interface{} {
	return attr.Value
}
func (attr XAttributeLiteral) GetKey() string {
	return attr.Key
}

func NewXAttributeLiteral(key string, value string, extensions []XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
