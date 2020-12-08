package xpath

import "sync"

/*
parserPreallocMem initializes all the memory pools required by the semantic function of the parser.
*/

var reductionPool = &sync.Pool{
	New: func() interface{} {
		return new(reductionImpl)
	},
}

var parserElementsPools []*pool

func parserPreallocMem(inputSize int, numThreads int) {
	parserElementsPools = make([]*pool, numThreads)

	poolSizePerThread := int(10000)
	constructor := func() interface{} {
		return new(element)
	}

	for i := 0; i < numThreads; i++ {
		parserElementsPools[i] = newPool(poolSizePerThread, constructor)
	}
}

/*
function is the semantic function of the parser
*/
func function(thread int, ruleNum uint16, lhs *symbol, rhs []*symbol) {
	switch ruleNum {
	case 0:
		NEW_AXIOM0 := lhs
		ELEM1 := rhs[0]

		NEW_AXIOM0.Child = ELEM1

		{
			NEW_AXIOM0.Value = ELEM1.Value
		}
	case 1:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		OPENCLOSETAG2 := rhs[1]

		ELEM0.Child = ELEM1
		ELEM1.Next = OPENCLOSETAG2

		{
			openCloseTag := OPENCLOSETAG2.Value.(openCloseTagSemanticValue)

			element := parserElementsPools[thread].Get().(*element)
			element.setFromSingleTag(openCloseTag)

			generativeNonTerminal := ELEM1.Value.(nonTerminal)
			reducedNonTerminal := newNonTerminal().setNode(element).setDirectChildAndInheritItsChildren(generativeNonTerminal)
			//logger.Printf("%v -> %v <%s />", reducedNonTerminal, generativeNonTerminal, element.name)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, generativeNonTerminal, nil)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 2:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		OPENTAG2 := rhs[1]
		ELEM3 := rhs[2]
		CLOSETAG4 := rhs[3]

		ELEM0.Child = ELEM1
		ELEM1.Next = OPENTAG2
		OPENTAG2.Next = ELEM3
		ELEM3.Next = CLOSETAG4

		{
			openTag := OPENTAG2.Value.(openTagSemanticValue)
			closeTag := CLOSETAG4.Value.(closeTagSemanticValue)

			element := parserElementsPools[thread].Get().(*element)
			element.setFromExtremeTags(openTag, closeTag)

			generativeNonTerminal := ELEM1.Value.(nonTerminal)
			wrappedNonTerminal := ELEM3.Value.(nonTerminal)
			reducedNonTerminal := newNonTerminal().setNode(element).setDirectChildAndInheritItsChildren(generativeNonTerminal)
			//logger.Printf("%v -> <%s> %v </%s>", reducedNonTerminal, element.name, wrappedNonTerminal, element.name)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, generativeNonTerminal, wrappedNonTerminal)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 3:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		OPENTAG2 := rhs[1]
		CLOSETAG3 := rhs[2]

		ELEM0.Child = ELEM1
		ELEM1.Next = OPENTAG2
		OPENTAG2.Next = CLOSETAG3

		{
			openTag := OPENTAG2.Value.(openTagSemanticValue)
			closeTag := CLOSETAG3.Value.(closeTagSemanticValue)

			element := parserElementsPools[thread].Get().(*element)
			element.setFromExtremeTags(openTag, closeTag)

			generativeNonTerminal := ELEM1.Value.(nonTerminal)
			reducedNonTerminal := newNonTerminal().setNode(element).setDirectChildAndInheritItsChildren(generativeNonTerminal)
			//logger.Printf("%v -> %v <%s></%s>", reducedNonTerminal, generativeNonTerminal, element.name, element.name)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, generativeNonTerminal, nil)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 4:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		TEXT2 := rhs[1]

		ELEM0.Child = ELEM1
		ELEM1.Next = TEXT2

		{
			tsv := TEXT2.Value.(textSemanticValue)

			text := new(text)
			text.setFromText(tsv)

			generativeNonTerminal := ELEM1.Value.(nonTerminal)

			reducedNonTerminal := newNonTerminal().setNode(text).setDirectChildAndInheritItsChildren(generativeNonTerminal)
			//logger.Printf("%v -> %v %s", reducedNonTerminal, generativeNonTerminal, text.data)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, generativeNonTerminal, nil)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 5:
		ELEM0 := lhs
		OPENCLOSETAG1 := rhs[0]

		ELEM0.Child = OPENCLOSETAG1

		{
			openCloseTag := OPENCLOSETAG1.Value.(openCloseTagSemanticValue)

			element := parserElementsPools[thread].Get().(*element)
			element.setFromSingleTag(openCloseTag)

			reducedNonTerminal := newNonTerminal().setNode(element)
			//logger.Printf("%v -> <%s />", reducedNonTerminal, element.name)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, nil, nil)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 6:
		ELEM0 := lhs
		OPENTAG1 := rhs[0]
		ELEM2 := rhs[1]
		CLOSETAG3 := rhs[2]

		ELEM0.Child = OPENTAG1
		OPENTAG1.Next = ELEM2
		ELEM2.Next = CLOSETAG3

		{
			openTag := OPENTAG1.Value.(openTagSemanticValue)
			closeTag := CLOSETAG3.Value.(closeTagSemanticValue)

			element := parserElementsPools[thread].Get().(*element)
			element.setFromExtremeTags(openTag, closeTag)

			wrappedNonTerminal := ELEM2.Value.(nonTerminal)
			reducedNonTerminal := newNonTerminal().setNode(element)
			//logger.Printf("%v -> <%s> %v </%s>", reducedNonTerminal, element.name, wrappedNonTerminal, element.name)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, nil, wrappedNonTerminal)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 7:
		ELEM0 := lhs
		OPENTAG1 := rhs[0]
		CLOSETAG2 := rhs[1]

		ELEM0.Child = OPENTAG1
		OPENTAG1.Next = CLOSETAG2

		{
			openTag := OPENTAG1.Value.(openTagSemanticValue)
			closeTag := CLOSETAG2.Value.(closeTagSemanticValue)

			element := parserElementsPools[thread].Get().(*element)
			element.setFromExtremeTags(openTag, closeTag)

			reducedNonTerminal := newNonTerminal().setNode(element)
			//logger.Printf("%v -> <%s></%s>", reducedNonTerminal, element.name, element.name)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, nil, nil)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal

		}
	case 8:
		ELEM0 := lhs
		TEXT1 := rhs[0]

		ELEM0.Child = TEXT1

		{
			tsv := TEXT1.Value.(textSemanticValue)

			text := new(text)
			text.setFromText(tsv)

			reducedNonTerminal := newNonTerminal().setNode(text)
			//logger.Printf("%v -> %s", reducedNonTerminal, text.data)

			reduction := reductionPool.Get().(reduction)
			reduction.setup(reducedNonTerminal, nil, nil)
			reduction.handle()

			reductionPool.Put(reduction)

			ELEM0.Value = reducedNonTerminal
		}
	}
}
