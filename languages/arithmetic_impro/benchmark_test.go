package arithmetic_impro

import (
	"fmt"
	"testing"
)

const (
	SMALL  = 1 + 2*3*(4+5)
	ONE_MB = (1*2*3 + 11*222*3333*(1+2)) * 25966
	TEN_MB = (1*2*3 + 11*222*3333*(1+2)) * 257473
)

func parseTestFile(b *testing.B, fileName string, thread_n int, exp_res int64) {
	root, err := ParseFile(fileName, thread_n)
	if err != nil {
		b.Errorf("unexpected error: %v", err)
	} else {
		res := *root.Value.(*int64)
		fmt.Printf("Analisys of file %s gave result: %d\n", fileName, res)
		if res != exp_res {
			b.Errorf("True res: %d\nCalculated res: %d\n", exp_res, res)
		}
	}

}

func BenchmarkSmall(b *testing.B) {
	parseTestFile(b, "data/small.txt", 1, SMALL)
}

func Benchmark1MB(b *testing.B) {
	parseTestFile(b, "data/1MB.txt", 1, ONE_MB)
}

func Benchmark10MB4T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 4, TEN_MB)
}

func Benchmark10MB1T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 1, TEN_MB)

}

func Benchmark10MB2T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 2, TEN_MB)

}

func Benchmark10MB8T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 8, TEN_MB)

}
