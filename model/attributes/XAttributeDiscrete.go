package attributes

type XAttributeDiscrete struct {
	XAttribute

	Value int

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeDiscrete) XAttributeMarker() {}

func NewXAttributeDiscrete(key string, value int, extensions []XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
