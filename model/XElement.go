package main

type XElement struct {
	XAttributable
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
	globalTraceAttrbutes  []XAttribute
	globalEventAttributes []XAttribute

	extensions  []XExtension
	classifiers []XClassifier
}

type XExtension struct {
	name   string
	uri    string
	prefix string

	logAttributes   []XAttribute
	eventAttributes []XAttribute
	metaAttributes  []XAttribute
}
type XClassifier struct {
	name string
	key  []string
}
