package xpath

import "fmt"

type node interface {
	position() *position
}

type attribute struct {
	key, value string
}

func (a *attribute) String() string {
	return fmt.Sprintf("%v=%v", a.key, a.value)
}

func newAttribute(key, value string) *attribute {
	return &attribute{key, value}
}

type element struct {
	name          string
	attributes    []*attribute
	posInDocument *position
}

func newElement(name string, attributes []*attribute, posInDocument *position) *element {
	return &element{name, attributes, posInDocument}
}

func (e *element) position() *position {
	return e.posInDocument
}

func (e *element) String() string {
	return fmt.Sprintf("<%v %v></%v>", e.name, e.attributes, e.name)
}

func (e *element) setFromExtremeTags(openTag openTagSemanticValue, closeTag closeTagSemanticValue) {
	if openTag.id != closeTag.id {
		panic("Invalid element construction")
	}
	e.name = openTag.id
	e.attributes = openTag.attributes
	e.posInDocument = newPosition(openTag.posInDocument.start, closeTag.posInDocument.end)
}

func (e *element) setFromSingleTag(openCloseTag openCloseTagSemanticValue) {
	e.name = openCloseTag.id
	e.attributes = openCloseTag.attributes
	e.posInDocument = openCloseTag.posInDocument
}

//Text node
type text struct {
	data          string
	posInDocument *position
}

func newText(data string, posInDocument *position) *text {
	return &text{data, posInDocument}
}

func (t *text) String() string {
	return fmt.Sprintf("text(%q)", t.data)
}

func (t *text) setFromText(tsv textSemanticValue) {
	t.data = tsv.data
	t.posInDocument = tsv.posInDocument
}

func (t *text) position() *position {
	return t.posInDocument
}
