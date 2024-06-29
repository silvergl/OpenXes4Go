package attributes

type XAttributeBoolean struct {
	XAttribute
	Value      bool
	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeBoolean) XAttributeMarker() {
}

func (attr XAttributeBoolean) GetExtensions() []XExtension {
	return attr.Extensions
}

func (attr XAttributeBoolean) GetAttributes() map[string]XAttribute {
	return attr.attributes
}

func (attr XAttributeBoolean) GetValue() interface{} {
	return attr.Value
}
func (attr XAttributeBoolean) GetKey() string {
	return attr.Key
}

/*
 * New XAttributeBoolean
 */
func NewXAttributeBoolean(key string, value bool, extensions []XExtension) (*XAttributeBoolean, error) {
	attr := &XAttributeBoolean{
		Key:        key,
		Extensions: extensions,
		Value:      value,
	}
	return attr, nil

}
