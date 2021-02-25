package main

import (
	"github.com/filgor84/gogopapageno/generator"
)

func main() {
	//generator.Generate("languages/arithmetic/lexer/arith.l", "languages/arithmetic/parser/arith.g", "languages/arithmetic")
	generator.Generate("languages/arithmetic_easy/lexer/arith.l", "languages/arithmetic_easy/parser/arith.g", "languages/arithmetic_easy")
}
