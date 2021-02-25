package main_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/filgor84/gogopapageno/languages/xpath"
)

const QUERYSETS_DATA_DIR = "data"

func TestA1_1_worker(t *testing.T) {
	input, _ := ioutil.ReadFile("data/standard.xml")
	xpath.Execute("A1").AgainstString(input).WithNumberOfThreads(1).Go()
}

func BenchmarkExecutor(b *testing.B) {
	querysets := []string{"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "B1"}
	input, _ := ioutil.ReadFile("data/standard.xml")
	b.ResetTimer()
	for _, queryset := range querysets {
		benchmarkQueryset(b, input, queryset)
	}
}

func benchmarkQueryset(b *testing.B, input []byte, queryset string) {
	numbersOfWorkers := intRange(1, 65)

	b.ResetTimer()
	b.Run(queryset, func(b *testing.B) {
		for _, nOfWorkers := range numbersOfWorkers {
			benchmarkQuerysetWithNumberOfWorkers(b, input, queryset, nOfWorkers)
		}
	})
}

func benchmarkQuerysetWithNumberOfWorkers(b *testing.B, input []byte, queryset string, nOfWorkers int) {
	b.Run(fmt.Sprintf("%d-workers", nOfWorkers), func(b *testing.B) {
		b.ResetTimer()
		var result []xpath.Position
		for i := 0; i < b.N; i++ {
			result, _ = xpath.Execute(queryset).AgainstString(input).WithNumberOfThreads(nOfWorkers).Go()
		}
		b.Logf("number of matches: %d\n", len(result))
	})
}

func intRange(inclMin, exclMax int) []int {
	interval := exclMax - inclMin
	result := make([]int, 0, interval)
	for i := inclMin; i < exclMax; i++ {
		result = append(result, i)
	}
	return result
}

func querysetDataFilePath(name string) string {
	fileName := strings.Join([]string{name, "xml"}, ".")
	return filepath.Join(QUERYSETS_DATA_DIR, fileName)
}
