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

func (attr XAttributeContinuos) GetAttributes() map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeContinuos) GetValue() interface{} {
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
