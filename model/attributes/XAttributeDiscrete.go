package attributes

type XAttributeDiscrete struct {
	XAttribute

	Value int

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeDiscrete) XAttributeMarker() {
}

func (attr XAttributeDiscrete) GetExtensions() []XExtension {
	return attr.Extensions
}

func (attr XAttributeDiscrete) GetExtension(name string) map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeDiscrete) GetValue() int {
	return attr.Value
}
func (attr XAttributeDiscrete) GetKey() string {
	return attr.Key
}

func NewXAttributeDiscrete(key string, value int, extensions []XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
