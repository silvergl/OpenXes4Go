package attributes

type XAttributeTimeStamp struct {
	XAttribute

	value string
}

func NewXAttributeDTimeStamp(key string, value string, extensions []XExtension) (*XAttributeTimeStamp, error) {
	attr := &XAttributeTimeStamp{
		XAttribute: XAttribute{
			key:        key,
			extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
