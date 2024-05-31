package model

import (
	"encoding/xml"
	"fmt"
)

type XEvent struct {
	XElement
}

func (x XEvent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//type xlog XLog
	start.Name = xml.Name{Local: "event"}

	// Marshal the MaleFriends with gender attribute
	for _, attr := range x.XElement.XAttributable.Attributes {
		attrElem := xml.StartElement{Name: xml.Name{Local: fmt.Sprintf("%T\n key", attr.Key)}}
		attrElem.Attr = []xml.Attr{{Name: xml.Name{Local: "key"}, Value: attr.Key}}

		if err := e.EncodeToken(attrElem); err != nil {
			return err
		}

	}

	// End the Person element
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
