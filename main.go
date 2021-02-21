package main

import (
	"github.com/filgor84/gogopapageno/generator"
)

func main() {
	//generator.Generate("languages/arithmetic/lexer/arith.l", "languages/arithmetic/parser/arith.g", "languages/arithmetic")
	generator.Generate("languages/xml/lexer/xml.l", "languages/xml/parser/xml.g", "languages/xml")
}
