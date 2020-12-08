
/*
parserPreallocMem initializes all the memory pools required by the semantic function of the parser.
*/

func parserPreallocMem(inputSize int, numThreads int) {
}

%%
%axiom ELEM
%%

ELEM : ELEM OPENTAG ELEM CLOSETAG
{  
    openTag := $2.Value.(openTagSemanticValue)
    closeTag := $4.Value.(closeTagSemanticValue)

    element := new(element)
    element.setFromExtremeTags(openTag, closeTag)
    
    generativeNonTerminal := $1.Value.(nonTerminal)
    wrappedNonTerminal := $3.Value.(nonTerminal)
    reducedNonTerminal := newNonTerminal().setNode(element).setDirectChildAndInheritItsChildren(generativeNonTerminal)
    logger.Printf("%v -> <%s> %v </%s>", reducedNonTerminal, element.name, wrappedNonTerminal, element.name)

    reduction := newReduction(reducedNonTerminal, generativeNonTerminal, wrappedNonTerminal)
    reduction.handle()

    $$.Value = reducedNonTerminal

} | OPENTAG ELEM CLOSETAG 
{
    openTag := $1.Value.(openTagSemanticValue)
    closeTag := $3.Value.(closeTagSemanticValue)

    element := new(element)
    element.setFromExtremeTags(openTag, closeTag)

    wrappedNonTerminal := $2.Value.(nonTerminal)
    reducedNonTerminal := newNonTerminal().setNode(element)
    logger.Printf("%v -> <%s> %v </%s>", reducedNonTerminal, element.name, wrappedNonTerminal, element.name)

    reduction := newReduction(reducedNonTerminal, nil, wrappedNonTerminal)
    reduction.handle()

    $$.Value = reducedNonTerminal

} | ELEM OPENTAG CLOSETAG 
{
    openTag := $2.Value.(openTagSemanticValue)
    closeTag := $3.Value.(closeTagSemanticValue)

    element := new(element)
    element.setFromExtremeTags(openTag, closeTag)
    
    generativeNonTerminal := $1.Value.(nonTerminal)
    reducedNonTerminal := newNonTerminal().setNode(element).setDirectChildAndInheritItsChildren(generativeNonTerminal)
    logger.Printf("%v -> %v <%s></%s>", reducedNonTerminal, generativeNonTerminal, element.name, element.name)

    reduction := newReduction(reducedNonTerminal, generativeNonTerminal, nil)
    reduction.handle()

    $$.Value = reducedNonTerminal

} | OPENTAG CLOSETAG
{
    openTag := $1.Value.(openTagSemanticValue)
    closeTag := $2.Value.(closeTagSemanticValue)

    element := new(element)
    element.setFromExtremeTags(openTag, closeTag)

    reducedNonTerminal := newNonTerminal().setNode(element)
    logger.Printf("%v -> <%s></%s>", reducedNonTerminal, element.name, element.name)

    reduction := newReduction(reducedNonTerminal, nil, nil)
    reduction.handle()

    $$.Value = reducedNonTerminal

} | ELEM OPENCLOSETAG 
{
    openCloseTag := $2.Value.(openCloseTagSemanticValue)

    element := new(element)
    element.setFromSingleTag(openCloseTag)

    generativeNonTerminal := $1.Value.(nonTerminal)
    reducedNonTerminal := newNonTerminal().setNode(element).setDirectChildAndInheritItsChildren(generativeNonTerminal)
    logger.Printf("%v -> %v <%s />", reducedNonTerminal, generativeNonTerminal, element.name)

    reduction := newReduction(reducedNonTerminal, generativeNonTerminal, nil)
    reduction.handle()

    $$.Value = reducedNonTerminal

} | OPENCLOSETAG  
{
    openCloseTag := $1.Value.(openCloseTagSemanticValue)

    element := new(element)
    element.setFromSingleTag(openCloseTag)

    reducedNonTerminal := newNonTerminal().setNode(element)
    logger.Printf("%v -> <%s />", reducedNonTerminal, element.name)
    
    reduction := newReduction(reducedNonTerminal, nil, nil)
    reduction.handle()

     $$.Value = reducedNonTerminal

} | ELEM TEXT 
{
    tsv := $2.Value.(textSemanticValue)

    text := new(text)
    text.setFromText(tsv)  

    generativeNonTerminal := $1.Value.(nonTerminal)

    reducedNonTerminal := newNonTerminal().setNode(text).setDirectChildAndInheritItsChildren(generativeNonTerminal)
    logger.Printf("%v -> %v %s", reducedNonTerminal, generativeNonTerminal, text.data)   

    reduction := newReduction(reducedNonTerminal, generativeNonTerminal, nil)  
    reduction.handle()

    $$.Value = reducedNonTerminal

} | TEXT 
{
    tsv := $1.Value.(textSemanticValue)

    text := new(text)
    text.setFromText(tsv)  

    reducedNonTerminal := newNonTerminal().setNode(text)
    logger.Printf("%v -> %s", reducedNonTerminal, text.data)   

    reduction := newReduction(reducedNonTerminal, nil, nil)  
    reduction.handle()

    $$.Value = reducedNonTerminal
};