package attributes

type XAttributeBoolean struct {
	XAttribute
	Value      bool
	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

func (attr XAttributeBoolean) XAttributeMarker() {}

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
