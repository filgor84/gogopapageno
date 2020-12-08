package xpath

import "fmt"

//nonTerminal represents a unique non terminal inside the syntax tree representing
//the XML document
type nonTerminal interface {
	setExecutionTable(exexTab executionTable) nonTerminal
	setNode(n interface{}) nonTerminal
	children() []nonTerminal
	setDirectChildAndInheritItsChildren(nonTerminal) nonTerminal
	executionTable() executionTable
	node() interface{}
	position() Position
}

func newNonTerminal() nonTerminal {
	return &nonTerminalImpl{}
}

type nonTerminalImpl struct {
	n       interface{}
	ch      []nonTerminal
	execTab executionTable
}

func (nt *nonTerminalImpl) String() string {
	if nt == nil {
		return "-"
	}
	return fmt.Sprintf("E(%p)", nt)
}

func (nt *nonTerminalImpl) setExecutionTable(executionTable executionTable) nonTerminal {
	nt.execTab = executionTable
	return nt
}

func (nt *nonTerminalImpl) executionTable() executionTable {
	return nt.execTab
}

func (nt *nonTerminalImpl) setNode(n interface{}) nonTerminal {
	nt.n = n
	return nt
}

func (nt *nonTerminalImpl) node() interface{} {
	return nt.n
}

func (nt *nonTerminalImpl) setDirectChildAndInheritItsChildren(child nonTerminal) nonTerminal {
	nt.ch = append(child.children(), child)
	return nt
}

func (nt *nonTerminalImpl) children() []nonTerminal {
	return nt.ch
}

func (nt *nonTerminalImpl) position() Position {
	if element, isElement := nt.n.(*element); isElement {
		return element.position()
	}

	if text, isText := nt.n.(*text); isText {
		return text.position()
	}

	return nil
}

//Position represents the position of some information inside a document
//in terms of number of characters from the beginning of the document.
type Position interface {
	Extremes() (start, end int)
	Start() int
	End() int
}

type position struct {
	start, end int
}

func newPosition(start, end int) *position {
	return &position{start, end}
}

func (p *position) String() string {
	return fmt.Sprintf("(%d , %d)", p.start, p.end)
}

func (p *position) Start() int {
	return p.start
}

func (p *position) End() int {
	return p.end
}

func (p *position) Extremes() (start, end int) {
	start = p.start
	end = p.end
	return
}

type openTagSemanticValue struct {
	id            string
	attributes    []*attribute
	posInDocument *position
}

func newOpenTagSemanticValue(id string, attributes []*attribute, posInDocument *position) *openTagSemanticValue {
	return &openTagSemanticValue{
		id:            id,
		attributes:    attributes,
		posInDocument: posInDocument,
	}
}

type closeTagSemanticValue struct {
	id            string
	posInDocument *position
}

func newCloseTagSemanticValue(id string, posInDocument *position) *closeTagSemanticValue {
	return &closeTagSemanticValue{
		id:            id,
		posInDocument: posInDocument,
	}
}

type openCloseTagSemanticValue struct {
	*openTagSemanticValue
}

func newOpenCloseTagSemanticValue(id string, attributes []*attribute, posInDocument *position) *openCloseTagSemanticValue {
	return &openCloseTagSemanticValue{newOpenTagSemanticValue(id, attributes, posInDocument)}
}

type textSemanticValue struct {
	data          string
	posInDocument *position
}

func newTextSemanticValue(data string, posInDocument *position) *textSemanticValue {
	return &textSemanticValue{
		data:          data,
		posInDocument: posInDocument,
	}
}
