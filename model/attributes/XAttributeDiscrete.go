package attributes

type XAttributeDiscrete struct {
	XAttribute

	value int
}

func NewXAttributeDiscrete(key string, value int, extensions []XExtension) (*XAttributeDiscrete, error) {
	attr := &XAttributeDiscrete{
		XAttribute: XAttribute{
			key:        key,
			extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
