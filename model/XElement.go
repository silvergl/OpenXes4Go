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

// MarshalXML customizes the XML marshaling for the Person struct
func (x XLog) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type xlog XLog
	start.Name = xml.Name{Local: "log"}
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xes.version"}, Value: x.Version},
		{Name: xml.Name{Local: "xes.features"}, Value: x.Features},
		{Name: xml.Name{Local: "openxes.version"}, Value: x.OpenXesVersion},
		{Name: xml.Name{Local: "xes.version"}, Value: x.Xmlns},
	}

	// Create the starting element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	startElement := xml.StartElement{Name: xml.Name{Local: "Global"}}
	startElement.Attr = []xml.Attr{{Name: xml.Name{Local: "scope"}, Value: "trace"}}

	if err := e.EncodeToken(startElement); err != nil {
		return err
	}
	// Marshal the MaleFriends with gender attribute
	for _, globalEvent := range x.GlobalEventAttributes {

		if err := e.EncodeElement(globalEvent.Key, xml.StartElement{Name: xml.Name{Local: "name"}}); err != nil {
			return err
		}

		if err := e.EncodeToken(xml.EndElement{Name: startElement.Name}); err != nil {
			return err
		}
	}
	eventElement := xml.StartElement{Name: xml.Name{Local: "Global"}}
	eventElement.Attr = []xml.Attr{{Name: xml.Name{Local: "scope"}, Value: "event"}}
	if err := e.EncodeToken(eventElement); err != nil {
		return err
	}
	// Marshal the FemaleFriends with gender attribute
	for _, globalEvent := range x.GlobalTraceAttributes {

		if err := e.EncodeToken(startElement); err != nil {
			return err
		}
		if err := e.EncodeElement(globalEvent.Key, xml.StartElement{Name: xml.Name{Local: "name"}}); err != nil {
			return err
		}

		if err := e.EncodeToken(xml.EndElement{Name: startElement.Name}); err != nil {
			return err
		}
	}

	// End the Person element
	return e.EncodeToken(xml.EndElement{Name: start.Name})
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
