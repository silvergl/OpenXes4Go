package model

import (
	attributes2 "github.com/openxes4go/model/attributes"
)

type XElement struct {
	attributes2.XAttributable
}

type XEvent struct {
	XElement
}

type XTrace struct {
	XElement
	events []XEvent
}

type XLog struct {
	XElement
	traces                []XTrace
	globalTraceAttributes []attributes2.XAttribute
	globalEventAttributes []attributes2.XAttribute

	extensions  []XExtension
	classifiers []XClassifier
}

type XExtension struct {
	name   string
	uri    string
	prefix string

	logAttributes   []attributes2.XAttribute
	eventAttributes []attributes2.XAttribute
	metaAttributes  []attributes2.XAttribute
}
type XClassifier struct {
	name string
	key  []string
}
