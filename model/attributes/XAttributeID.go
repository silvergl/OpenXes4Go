package attributes

type XAttributeID struct {
	XAttribute

	Value string //TODO ID OWN TYPE

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func NewXAttributeID(key string, value string, extensions []XExtension) (*XAttributeID, error) {
	attr := &XAttributeID{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
