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
}

type XClassifier struct {
	name string
	key  []string
}
