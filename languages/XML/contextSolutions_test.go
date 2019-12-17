package XML

import (
	"reflect"
	"testing"
)

func TestContextSolutionsMap(t *testing.T) {
	t.Run(`hasSolutions(ctx)`, func(t *testing.T) {
		t.Run(`hasSolutionsFor(ctx)=false`, func(t *testing.T) {
			contextSolutionsMap := newContexSolutionsMap()

			context1 := newNonTerminal()
			solution1 := newNonTerminal()

			context2 := newNonTerminal()

			contextSolutionsMap.addContextSolution(context1, solution1)

			if contextSolutionsMap.hasSolutionsFor(context2) {
				t.Error(`hasSolutions(ctx)=true | want false`)
			}
		})

		t.Run(`hasSolutionsFor(ctx)=true`, func(t *testing.T) {
			contextSolutionsMap := newContexSolutionsMap()
			context := newNonTerminal()
			solution := newNonTerminal()

			contextSolutionsMap.addContextSolution(context, solution)

			if !contextSolutionsMap.hasSolutionsFor(context) {
				t.Error(`hasSolutions(ctx)=false | want true`)
			}
		})
	})

	t.Run(`solutionsFor(ctx, maps)`, func(t *testing.T) {

		root := new(nonTerminalImpl)
		solution1 := new(nonTerminalImpl)
		solution2 := new(nonTerminalImpl)

		firstStepFromRootMap := newContexSolutionsMap()
		firstStepFromRootMap.addContextSolution(root, solution1)
		firstStepFromRootMap.addContextSolution(root, solution2)

		solution11 := new(nonTerminalImpl)
		solution21 := new(nonTerminalImpl)
		solution22 := new(nonTerminalImpl)

		secondStepFromRootMap := newContexSolutionsMap()
		secondStepFromRootMap.addContextSolution(solution1, solution11)
		secondStepFromRootMap.addContextSolution(solution2, solution21)
		secondStepFromRootMap.addContextSolution(solution2, solution22)

		solution221 := new(nonTerminalImpl)
		thirdStepFromRootMap := newContexSolutionsMap()
		thirdStepFromRootMap.addContextSolution(solution22, solution221)

		var tests = []struct {
			context                       nonTerminal
			startContextSolutionMap       contextSolutionsMap
			subsequentContextSolutionMaps []contextSolutionsMap
			want                          []nonTerminal
		}{
			{
				context:                       root,
				startContextSolutionMap:       firstStepFromRootMap,
				subsequentContextSolutionMaps: nil,
				want: []nonTerminal{solution1, solution2},
			},
			{
				context:                       root,
				startContextSolutionMap:       firstStepFromRootMap,
				subsequentContextSolutionMaps: []contextSolutionsMap{secondStepFromRootMap},
				want: []nonTerminal{solution11, solution21, solution22},
			},
			{
				context:                       root,
				startContextSolutionMap:       firstStepFromRootMap,
				subsequentContextSolutionMaps: []contextSolutionsMap{secondStepFromRootMap, thirdStepFromRootMap},
				want: []nonTerminal{solution221},
			},
		}

		for _, test := range tests {
			got := test.startContextSolutionMap.solutionsFor(test.context, test.subsequentContextSolutionMaps...)

			if !areSolutionsEqual(got, test.want) {
				t.Errorf(`solutionsFor(%p, %v)=%v | want %v`, test.context, test.subsequentContextSolutionMaps, got, test.want)
			}
		}

	})

	t.Run(`transitiveClosure(maps)`, func(t *testing.T) {

		root1 := new(nonTerminalImpl)
		solution11 := new(nonTerminalImpl)

		root2 := new(nonTerminalImpl)
		solution21 := new(nonTerminalImpl)
		solution22 := new(nonTerminalImpl)

		firstStepFromRootsMap := newContexSolutionsMap()
		firstStepFromRootsMap.addContextSolution(root1, solution11)
		firstStepFromRootsMap.addContextSolution(root2, solution21, solution22)

		solution111 := new(nonTerminalImpl)

		solution211 := new(nonTerminalImpl)
		solution212 := new(nonTerminalImpl)

		secondStepFromRootsMap := newContexSolutionsMap()
		secondStepFromRootsMap.addContextSolution(solution11, solution111)
		secondStepFromRootsMap.addContextSolution(solution21, solution211, solution212)

		var tests = []struct {
			startContextSolutionMap       contextSolutionsMap
			subsequentContextSolutionMaps []contextSolutionsMap
			want                          contextSolutionsMap
		}{
			{
				startContextSolutionMap:       firstStepFromRootsMap,
				subsequentContextSolutionMaps: nil,
				want: &contextSolutionsMapImpl{
					m: map[nonTerminal][]nonTerminal{
						root1: []nonTerminal{solution11},
						root2: []nonTerminal{solution21, solution22},
					},
				},
			},
			{
				startContextSolutionMap:       firstStepFromRootsMap,
				subsequentContextSolutionMaps: []contextSolutionsMap{secondStepFromRootsMap},
				want: &contextSolutionsMapImpl{
					m: map[nonTerminal][]nonTerminal{
						root1: []nonTerminal{solution111},
						root2: []nonTerminal{solution211, solution212},
					},
				},
			},
		}

		for _, test := range tests {
			got := test.startContextSolutionMap.transitiveClosure(test.subsequentContextSolutionMaps...)

			if !reflect.DeepEqual(got, test.want) {
				t.Error(`transitiveClosure does NOT return correctly`)
			}
		}
	})

	t.Run(`merge(incoming)`, func(t *testing.T) {
		t.Run(`merge(nil)`, func(t *testing.T) {
			contextSolutionsMap := newContexSolutionsMap()

			if result, ok := contextSolutionsMap.merge(nil); result != contextSolutionsMap || ok {
				t.Errorf(`merge(nil) return incorrectly`)
			}
		})

		t.Run(`merge(incoming)`, func(t *testing.T) {
			receiverContextSolutionMap := newContexSolutionsMap()
			context1 := newNonTerminal()
			solution1 := newNonTerminal()
			receiverContextSolutionMap.addContextSolution(context1, solution1)

			incomingContextSolutionMap := newContexSolutionsMap()
			solution2 := newNonTerminal()
			incomingContextSolutionMap.addContextSolution(context1, solution2)

			if result, ok := receiverContextSolutionMap.merge(incomingContextSolutionMap); result != receiverContextSolutionMap || !ok {
				t.Error(`cannot merge incoming context solution map into receiver context solution map`)
			}

			implReceiverContextsolutionMap := receiverContextSolutionMap.(*contextSolutionsMapImpl)

			if len(implReceiverContextsolutionMap.m[context1]) != 2 {
				t.Error(`merged context solution map does does NOT have the right number of solutions for context`)
			}

		})
	})
}

//utils
func areSolutionsEqual(a, b []nonTerminal) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
