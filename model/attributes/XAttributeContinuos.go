package attributes

type XAttributeContinuos struct {
	XAttribute
	value float64
}

func NewXAttributeContinuos(key string, value float64, extensions []XExtension) (*XAttributeContinuos, error) {
	attr := &XAttributeContinuos{
		XAttribute: XAttribute{
			key:        key,
			extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
