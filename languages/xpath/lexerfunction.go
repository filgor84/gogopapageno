package xpath

import (
	"regexp"
	"strings"
)

//A regexp.Regexp is safe for concurrent use by multiple goroutines, except for
//consifugration methods, such as Longest.
var generalTagRegexp *(regexp.Regexp)
var attributesRegexp *(regexp.Regexp)

/*func getIdAndAttributesListFrom(text string)(id string, attributesList []*attribute){
	tagMatch := generalTagRegexp.FindStringSubmatch(text)

	id = tagMatch[1]
	attributesString := tagMatch[2]

	if(attributesString != ""){
		attributesMatches := attributesRegexp.FindAllStringSubmatch(attributesString, -1)
		for _, attributeMatch := range attributesMatches{
			attribute := newAttribute(attributeMatch[1], attributeMatch[2])
			attributesList = append(attributesList, attribute)
		}
	}
	return
}*/

func getIdAndAttributesListFrom(text string) (string, []*attribute) {
	if text[0] != '<' {
		return "", nil
	}

	leftmostLimit := 1

	if text[leftmostLimit] == '/' {
		leftmostLimit++
	}

	rightmostLimit := len(text) - 1
	if text[rightmostLimit] != '>' {
		return "", nil
	}

	if text[rightmostLimit-1] == '/' {
		rightmostLimit--
	}

	textToParse := text[leftmostLimit:rightmostLimit]
	firstSpace := strings.IndexRune(textToParse, ' ')

	if firstSpace == -1 {
		return textToParse, nil
	}

	attributesSpace := textToParse[firstSpace+1:]
	rawAttributes := strings.Split(attributesSpace, " ")
	attributesList := make([]*attribute, 0, len(rawAttributes))

	for _, rawAttribute := range rawAttributes {
		equalPosition := strings.IndexRune(rawAttribute, '=')
		attributesList = append(attributesList, &attribute{rawAttribute[:equalPosition], rawAttribute[equalPosition+1:]})
	}

	return textToParse[:firstSpace], attributesList
}

func lexerPreallocMem(inputSize int, numThreads int) {
	generalTagRegexp = regexp.MustCompile(`^<\/?([a-zA-Z0-9_\-:]+)(?:\s*)([^\/>]*)\/?>$`)
	attributesRegexp = regexp.MustCompile(`([a-zA-Z0-9_\-:]+)(?:=")([^"]+)(?:")`)
}

/*
lexerFunction is the semantic function of the lexer.
*/
func lexerFunction(thread int, ruleNum int, yytext string, absPos *position, genSym *symbol) int {
	switch ruleNum {
	case 0:
		{
			id, attributesList := getIdAndAttributesListFrom(yytext)
			semanticValue := newOpenTagSemanticValue(id, attributesList, absPos)

			*genSym = symbol{OPENTAG, 0, *semanticValue, nil, nil}
			return _LEX_CORRECT
		}
	case 1:
		{
			id, attributesList := getIdAndAttributesListFrom(yytext)
			semanticValue := newOpenCloseTagSemanticValue(id, attributesList, absPos)

			*genSym = symbol{OPENCLOSETAG, 0, *semanticValue, nil, nil}
			return _LEX_CORRECT
		}
	case 2:
		{
			id, _ := getIdAndAttributesListFrom(yytext)
			semanticValue := newCloseTagSemanticValue(id, absPos)

			*genSym = symbol{CLOSETAG, 0, *semanticValue, nil, nil}
			return _LEX_CORRECT
		}
	case 3:
		{
			return _SKIP
		}
	case 4:
		{
			return _SKIP
		}
	case 5:
		{
			return _SKIP
		}
	case 6:
		{
			semanticValue := newTextSemanticValue(yytext, absPos)

			*genSym = symbol{TEXT, 0, *semanticValue, nil, nil}
			return _LEX_CORRECT
		}
	case 7:
		{
			return _ERROR
		}
	}
	return _ERROR
}
