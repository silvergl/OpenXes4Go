package attributes

type XAttributeBoolean struct {
	XAttribute
	value bool
}

/*
 * New XAttributeBoolean
 */
func NewXAttributeBoolean(key string, value bool, extensions []XExtension) (*XAttributeBoolean, error) {
	attr := &XAttributeBoolean{
		XAttribute: XAttribute{
			key:        key,
			extensions: extensions,
		},
		value: value,
	}
	return attr, nil

}
