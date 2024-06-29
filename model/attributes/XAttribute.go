package attributes

type XAttribute interface {
	XAttributeMarker()
	GetKey() string
	GetValue() interface{}
	GetExtensions() interface{}
	GetAttributes() interface{}
	//Attributable
}
