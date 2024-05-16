package attributes

type XAttributeID struct {
	XAttribute

	value string //TODO ID OWN TYPE
}

func NewXAttributeID(key string, value string, extensions []XExtension) (*XAttributeID, error) {
	attr := &XAttributeID{
		XAttribute: XAttribute{
			key:        key,
			extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
