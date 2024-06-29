package attributes

type XAttributeContinuos struct {
	XAttribute
	Value float64

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeContinuos) XAttributeMarker() {}

func NewXAttributeContinuos(key string, value float64, extensions []XExtension) (*XAttributeContinuos, error) {
	attr := &XAttributeContinuos{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
