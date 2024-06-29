package extensions

import attributes2 "github.com/silvergl/OpenXes4Go/model/attributes"

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
