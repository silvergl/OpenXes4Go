package main

type XAttributable struct {
	attributes []XAttribute
}

type XAttribute struct {
	XAttributable
}

type XAttributeBoolean struct {
	XAttribute
}

type XAttributeContainer struct {
	XAttribute
}

type XAttributeContinuos struct {
	XAttribute
}

type XAttributeDiscrete struct {
	XAttribute
}

type XAttributeID struct {
	XAttribute
}

type XAttributeList struct {
	XAttribute
}

type XAttributeLiteral struct {
	XAttribute
}

type XAttributeTimestamp struct {
	XAttribute
}
