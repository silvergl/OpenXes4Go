package model

type XTrace struct {
	XElement
	events []XEvent `xml:"events>event"`
}
