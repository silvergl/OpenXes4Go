package model

import (
	"encoding/xml"
	"fmt"
)

type XLog struct {
	XElement
	XMLName               xml.Name
	Version               string
	Features              string
	OpenXesVersion        string
	Xmlns                 string
	Traces                []XTrace
	GlobalTraceAttributes []Global
	GlobalEventAttributes []Global
	extensions            []XExtension
	classifiers           []XClassifier
}

// MarshalXML customizes the XML marshaling for the Person struct
func (x XLog) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//type xlog XLog
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
		attrElem := xml.StartElement{Name: xml.Name{Local: fmt.Sprintf("%T\n key", globalEvent.Key)}}
		attrElem.Attr = []xml.Attr{{Name: xml.Name{Local: "key"}, Value: globalEvent.Key}}

		if err := e.EncodeToken(attrElem); err != nil {
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
		attrElem := xml.StartElement{Name: xml.Name{Local: fmt.Sprintf("%T\n key", globalEvent.Key)}}
		attrElem.Attr = []xml.Attr{{Name: xml.Name{Local: "key"}, Value: globalEvent.Key}}
		if err := e.EncodeToken(startElement); err != nil {
			return err
		}

		if err := e.EncodeToken(xml.EndElement{Name: startElement.Name}); err != nil {
			return err
		}
	}

	// End the Person element
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
