package attributes

type XAttributeContinuos struct {
	XAttribute
	Value float64

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeContinuos) XAttributeMarker() {}

func (attr XAttributeContinuos) GetExtensions() []XExtension {
	return attr.Extensions
}

func (attr XAttributeContinuos) GetExtension(name string) map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeContinuos) GetValue() float64 {
	return attr.Value
}
func (attr XAttributeContinuos) GetKey() string {
	return attr.Key
}
func NewXAttributeContinuos(key string, value float64, extensions []XExtension) (*XAttributeContinuos, error) {
	attr := &XAttributeContinuos{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
