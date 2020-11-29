%cut <

IDENTIFIER [a-zA-Z0-9_\-:]+
INFO [^<]+
SPACE [ \t]
NEWLINE [\r\n]
VALUE ([^"\\]|\\.)*

%%

<{IDENTIFIER}({SPACE}+{IDENTIFIER}="{VALUE}")*>
{
	id, attributesList := getIdAndAttributesListFrom(yytext)
	semanticValue := newOpenTagSemanticValue(id, attributesList, absPos)

	*genSym = symbol{OPENTAG, 0, *semanticValue, nil, nil}
	return _LEX_CORRECT
}
<{IDENTIFIER}({SPACE}+{IDENTIFIER}="{VALUE}")*/>
{
	id, attributesList := getIdAndAttributesListFrom(yytext)
	semanticValue := newOpenCloseTagSemanticValue(id, attributesList, absPos)

	*genSym = symbol{OPENCLOSETAG, 0, *semanticValue, nil, nil}
	return _LEX_CORRECT
}
</{IDENTIFIER}>
{
	id, _ := getIdAndAttributesListFrom(yytext)
	semanticValue := newCloseTagSemanticValue(id, absPos)

	*genSym = symbol{CLOSETAG, 0, *semanticValue, nil, nil}
	return _LEX_CORRECT
}
({SPACE}|{NEWLINE})+
{
	return _SKIP
}
<?[^?]+?>
{
	return _SKIP
}
<![^>]+>
{
	return _SKIP
}
{INFO}
{
	semanticValue := newTextSemanticValue(yytext, absPos)

	*genSym = symbol{TEXT, 0, *semanticValue, nil, nil}
	return _LEX_CORRECT
}
.
{
	return _ERROR
}
%%

import (
	"regexp"
)

//A regexp.Regexp is safe for concurrent use by multiple goroutines, except for
//consifugration methods, such as Longest.
var generalTagRegexp *(regexp.Regexp)
var attributesRegexp *(regexp.Regexp)

func getIdAndAttributesListFrom(text string)(id string, attributesList []*attribute){
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
}

func lexerPreallocMem(inputSize int, numThreads int){
	generalTagRegexp = regexp.MustCompile(`^<\/?([a-zA-Z0-9_\-:]+)(?:\s*)([^\/>]*)\/?>$`)
	attributesRegexp = regexp.MustCompile(`([a-zA-Z0-9_\-:]+)(?:=")([^"]+)(?:")`)
}
