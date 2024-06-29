package attributes

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
