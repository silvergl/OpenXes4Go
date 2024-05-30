package model

import (
	"encoding/xml"
	attributes2 "github.com/openxes4go/model/attributes"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:"lastName"`
	Age       int      `xml:"age"`
}
type XElement struct {
	attributes2.XAttributable
}

type XEvent struct {
	XElement
}

type XTrace struct {
	XElement
	events []XEvent `xml:"events>event"`
}

type XLog struct {
	XElement
	XMLName               xml.Name `xml:"log"`
	Version               string   `xml:"xes.version,attr"`
	Features              string   `xml:"xes.features,attr"`
	OpenXesVersion        string   `xml:"openxes.version,attr"`
	Xmlns                 string   `xml:"xmlns,attr"`
	Traces                []XTrace `xml:"trace"`
	GlobalTraceAttributes []Global `xml:"global"`
	GlobalEventAttributes []Global `xml:"global"`

	extensions  []XExtension
	classifiers []XClassifier
}

type Global struct {
	Scope string `xml:"scope,attr"`
	attributes2.XAttribute
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
