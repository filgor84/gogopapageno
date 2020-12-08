package xpath

import (
	"fmt"
)

type reduction interface {
	setup(reducedNT, generativeNT, wrappedNT nonTerminal)
	handle()
	reset()
}

func (reduction *reductionImpl) setup(reducedNT, generativeNT, wrappedNT nonTerminal) {
	var updatingExecutionTable executionTable

	if wrappedNT != nil {
		updatingExecutionTable = wrappedNT.executionTable()
	} else {
		updatingExecutionTable = udpeGlobalTable.newExecutionTable()
	}

	reduction.reducedNT = reducedNT
	reduction.generativeNT = generativeNT
	reduction.wrappedNT = wrappedNT
	reduction.updatingExecutionTable = updatingExecutionTable
}

func (reduction *reductionImpl) reset() {
	reduction.reducedNT = nil
	reduction.generativeNT = nil
	reduction.wrappedNT = nil
	reduction.updatingExecutionTable = nil
	reduction.globalUdpeRecordBeingConsidered = nil
}

type reductionImpl struct {
	reducedNT, generativeNT, wrappedNT nonTerminal
	updatingExecutionTable             executionTable
	globalUdpeRecordBeingConsidered    globalUdpeRecord
}

func (reduction *reductionImpl) handle() {
	defer reduction.avoidMemoryLeaksAtTheEndOfHandling()

	reduction.iterateOverAllGlobalUdpeRecordsAndExecuteMainPhases()
	reduction.prepareUpdatingExecutionTableToBePropagatedToReducedNT()
	reduction.propagateUpdatingExecutionTableToReducedNT()
}

func (reduction *reductionImpl) avoidMemoryLeaksAtTheEndOfHandling() {
	reduction.reducedNT = nil
	reduction.generativeNT = nil
	reduction.wrappedNT = nil
	reduction.updatingExecutionTable = nil
	reduction.globalUdpeRecordBeingConsidered = nil
}

func (reduction *reductionImpl) iterateOverAllGlobalUdpeRecordsAndExecuteMainPhases() {
	udpeGlobalTable.iterate(func(id int, globalRecord globalUdpeRecord) {
		reduction.globalUdpeRecordBeingConsidered = globalRecord
		updatingExecutionRecord, err := reduction.updatingExecutionTable.recordByID(id)
		if err != nil {
			panic(fmt.Sprintf("cannot retrieve execution record for udpe with id: %d", id))
		}

		reduction.addNewExecutionThreadsToExecutionRecord(updatingExecutionRecord)
		updatingExecutionRecord.updateAllExecutionThreads(reduction.reducedNT)
		updatingExecutionRecord.stopUnfoundedSpeculativeExecutionThreads(reduction.updatingExecutionTable.evaluateID)
		updatingExecutionRecord.saveReducedNTAsContextOrSolutionlIntoCompletedExecutionThreads(reduction.reducedNT)
		updatingExecutionRecord.produceContextSolutionsOutOfCompletedNonSpeculativeExecutionThreads()
	})
}

func (reduction *reductionImpl) prepareUpdatingExecutionTableToBePropagatedToReducedNT() {
	if reduction.generativeNT != nil {
		reduction.mergeUpdatingExecutionTableWithUnchangedExecutionTable()
	}
}

func (reduction *reductionImpl) propagateUpdatingExecutionTableToReducedNT() {
	reduction.reducedNT.setExecutionTable(reduction.updatingExecutionTable)
}

func (reduction *reductionImpl) mergeUpdatingExecutionTableWithUnchangedExecutionTable() {
	unchangedExecutionTable := reduction.generativeNT.executionTable()
	_, ok := reduction.updatingExecutionTable.merge(unchangedExecutionTable)
	if !ok {
		panic(`reduction handle node error: can NOT merge execution tables`)
	}
}

//phase1 checks if the udpe's entry point matches current reduction, without updating the path pattern,
//and, if it maches, creates new execution threads accordingly
func (reduction *reductionImpl) addNewExecutionThreadsToExecutionRecord(executionRecord executionRecord) {
	udpe := reduction.globalUdpeRecordBeingConsidered.udpe()
	entryPoint := udpe.entryPoint()

	if _, _, ok := entryPoint.matchWithReductionOf(reduction.reducedNT.node(), false); !ok {
		return
	}

	switch udpeType := reduction.globalUdpeRecordBeingConsidered.udpeType(); udpeType {
	case FPE:
		executionRecord.addExecutionThread(nil, reduction.reducedNT, entryPoint)
	case RPE:
		if reduction.wrappedNT == nil {
			return
		}
		executionRecord.addExecutionThread(reduction.wrappedNT, nil, entryPoint)
		childrenOfWrappedNT := reduction.wrappedNT.children()
		for _, child := range childrenOfWrappedNT {
			executionRecord.addExecutionThread(child, nil, udpe.entryPoint())
		}

	default:
		panic(fmt.Sprintf(`adding new execution threads to execution record: unknown udpe type %q`, udpeType))
	}
}
