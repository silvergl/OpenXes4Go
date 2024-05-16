package attributes

type XAttribute struct {
	XAttributable

	key        string
	extensions []XExtension
	attributes map[string]XAttribute
}

/*
 * Empty Attribute
 */
func NewXAtributeEmpty(key string) (*XAttribute, error) {
	attr := &XAttribute{
		key: key,

		XAttributable: XAttributable{
			attributes: make([]XAttribute, 0),
		},
	}
	return attr, nil
}

/*
 * Creates New Attribute
 */
func NewXAtribute(key string, extensions []XExtension) (*XAttribute, error) {
	attr := &XAttribute{
		key:        "0",
		extensions: extensions,
		XAttributable: XAttributable{
			attributes: make([]XAttribute, 0),
		},
	}
	return attr, nil
}
