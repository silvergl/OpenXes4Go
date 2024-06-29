package attributes

type XAttribute interface {
	XAttributeMarker()
	GetKey() string
	GetValue() interface{}
	GetExtensions() []XExtension
	GetAttributes() map[string]XAttribute
	//Attributable
}
