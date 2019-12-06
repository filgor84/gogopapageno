GoGoPAPAGENO
========

GoGoPAPAGENO (PArallel PArser GENeratOr) is a parallel parser generator based on Floyd's Operator Precedence Grammars.

It generates parallel Go parsers starting from a lexer and a grammar specification.
These specification files resemble Flex and Bison ones, although with some differences.

The generated parsers are self-contained and can be used without further effort.

This fork of Gopapageno, by Simone Guidi, is supposed to solve some original performance issues and add support for further languages (XPath queries and others).

This work is based on [Papageno](https://github.com/PAPAGENO-devels/papageno), a C parallel parser generator.


### Authors and Contributors

 * Filippo Gorlero <filgor84@gmail.com>
 * Simone Guidi <simone.guidi@mail.polimi.it>
