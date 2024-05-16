package main

type XAttributeLiteral struct {
	XAttribute

	value string
}

func NewXAttributeLiteral(key string, value string, extensions []XExtension) (*XAttributeLiteral, error) {
	attr := &XAttributeLiteral{
		XAttribute: XAttribute{
			key:        key,
			extensions: extensions,
		},
		value: value,
	}

	return attr, nil
}
