package attributes

type XAttributeTimeStamp struct {
	XAttribute

	Value string

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeTimeStamp) XAttributeMarker() {
}

func (attr XAttributeTimeStamp) GetExtensions() []XExtension {
	return attr.Extensions
}

func (attr XAttributeTimeStamp) GetAttributes() map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeTimeStamp) GetValue() interface{} {
	return attr.Value
}
func (attr XAttributeTimeStamp) GetKey() string {
	return attr.Key
}
func NewXAttributeDTimeStamp(key string, value string, extensions []XExtension) (*XAttributeTimeStamp, error) {

	attr := &XAttributeTimeStamp{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}

	return attr, nil
}
