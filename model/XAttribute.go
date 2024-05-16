package model

type XAttribute struct {
	XAttributable

	Key        string
	Extensions []XExtension
	attributes map[string]XAttribute
}

/*
 * Empty Attribute
 */
func NewXAttributeEmpty(key string) (*XAttribute, error) {
	attr := &XAttribute{
		Key: key,

		XAttributable: XAttributable{
			attributes: make([]XAttribute, 0),
		},
	}
	return attr, nil
}

/*
 * Creates New Attribute
 */
func NewXAttribute(key string, extensions []XExtension) (*XAttribute, error) {
	attr := &XAttribute{
		Key:        "0",
		Extensions: extensions,
		XAttributable: XAttributable{
			attributes: make([]XAttribute, 0),
		},
	}
	return attr, nil
}
