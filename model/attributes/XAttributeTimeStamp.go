package attributes

type XAttributeTimeStamp struct {
	XAttribute

	value string

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func NewXAttributeDTimeStamp(key string, value string, extensions []XExtension) (*XAttributeTimeStamp, error) {

	attr := &XAttributeTimeStamp{
		Key:        key,
		Extensions: extensions,
		value:      value,
	}

	return attr, nil
}
