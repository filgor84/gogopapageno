package xpath

import (
	"errors"
)

type contextSolutionsMap interface {
	addContextSolution(ctx nonTerminal, sols ...nonTerminal)
	transitiveClosure(maps ...contextSolutionsMap) contextSolutionsMap
	hasSolutionsFor(ctx nonTerminal) bool
	solutionsFor(ctx nonTerminal, maps ...contextSolutionsMap) []nonTerminal
	merge(incoming contextSolutionsMap) (result contextSolutionsMap, ok bool)
	convertToGroupOfSolutionsPositions() []Position
}

type implementedMapType map[nonTerminal][]nonTerminal

type contextSolutionsMapImpl struct {
	m implementedMapType
}

func newContexSolutionsMap() contextSolutionsMap {
	return &contextSolutionsMapImpl{
		m: make(implementedMapType),
	}
}

func (ctxSolMap *contextSolutionsMapImpl) addContextSolution(ctx nonTerminal, sols ...nonTerminal) {
	ctxSolMap.m[ctx] = append(ctxSolMap.m[ctx], sols...)
}

func (ctxSolMap *contextSolutionsMapImpl) hasSolutionsFor(ctx nonTerminal) bool {
	return len(ctxSolMap.solutionsFor(ctx)) > 0
}

func (ctxSolMap *contextSolutionsMapImpl) solutionsFor(ctx nonTerminal, maps ...contextSolutionsMap) (solutions []nonTerminal) {
	solutions = ctxSolMap.m[ctx]

	for currentMapIdx := 0; currentMapIdx < len(maps); currentMapIdx++ {
		tmpNodesToVisit := []nonTerminal{}

		for len(solutions) > 0 {
			currentNode := solutions[0]
			solutions = solutions[1:]
			tmpNodesToVisit = append(tmpNodesToVisit, maps[currentMapIdx].solutionsFor(currentNode)...)
		}
		solutions = tmpNodesToVisit
	}
	return
}

func (ctxSolMap *contextSolutionsMapImpl) transitiveClosure(maps ...contextSolutionsMap) (result contextSolutionsMap) {
	result = newContexSolutionsMap()

	for context := range ctxSolMap.m {
		solutionsReachableFromContext := ctxSolMap.solutionsFor(context, maps...)
		result.addContextSolution(context, solutionsReachableFromContext...)
	}
	return
}

func (ctxSolMap *contextSolutionsMapImpl) convertToGroupOfSolutionsPositions() (positions []Position) {
	for _, solutions := range ctxSolMap.m {
		for _, solution := range solutions {
			positions = append(positions, solution.position())
		}
	}
	return
}

func transitiveClosure(maps []contextSolutionsMap) contextSolutionsMap {
	start := maps[0]
	return start.transitiveClosure(maps[1:]...)
}

func (ctxSolMap *contextSolutionsMapImpl) merge(incoming contextSolutionsMap) (result contextSolutionsMap, ok bool) {
	result = ctxSolMap
	if incoming == nil {
		return
	}

	incomingImpl, ok := incoming.(*contextSolutionsMapImpl)

	if !ok {
		return
	}

	for k, v := range incomingImpl.m {
		ctxSolMap.m[k] = append(ctxSolMap.m[k], v...)
	}
	ok = true
	return
}

//solutionsFor returns all the solutions that are reachable from the specified context
//by traversing all the contextSolutionsMaps which are passed as parameters
func solutionsFor(context nonTerminal, maps ...contextSolutionsMap) (solutions []nonTerminal, err error) {
	if context == nil {
		err = errors.New("context can NOT be nil")
		return
	}
	solutions = maps[0].solutionsFor(context)

	for currentMapIdx := 1; currentMapIdx < len(maps); currentMapIdx++ {
		tmpNodesToVisit := []nonTerminal{}

		for len(solutions) > 0 {
			currentNode := solutions[0]
			solutions = solutions[1:]
			tmpNodesToVisit = append(tmpNodesToVisit, maps[currentMapIdx].solutionsFor(currentNode)...)
		}
		solutions = tmpNodesToVisit
	}
	return
}
