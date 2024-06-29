package attributes

type XAttributeID struct {
	XAttribute

	Value string //TODO ID OWN TYPE

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeID) XAttributeMarker() {
}

func (attr XAttributeID) GetExtensions() []XExtension {
	return attr.Extensions
}

func (attr XAttributeID) GetExtension(name string) map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeID) GetValue() string {
	return attr.Value
}
func (attr XAttributeID) GetKey() string {
	return attr.Key
}

func NewXAttributeID(key string, value string, extensions []XExtension) (*XAttributeID, error) {
	attr := &XAttributeID{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
