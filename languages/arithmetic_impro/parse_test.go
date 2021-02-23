package arithmetic_impro

import (
	"testing"
)

const (
	SMALL  = 1 + 2*3*(4+5)
	ONE_MB = (1*2*3 + 11*222*3333*(1+2)) * 25966
	TEN_MB = (1*2*3 + 11*222*3333*(1+2)) * 257473
)

func parseTestFile(b *testing.B, fileName string, thread_n int, exp_res int64) {
	for i := 0; i < b.N; i++ {
		root, err := ParseFile(fileName, thread_n)
		if err != nil {
			b.Errorf("unexpected error: %v", err)
		} else {
			res := *root.Value.(*int64)
			if res != exp_res {
				b.Errorf("True res: %d\nExpected res: %d\n", exp_res, res)
			}
		}
	}

}

func BenchmarkSmall(b *testing.B) {
	parseTestFile(b, "data/small.txt", 1, SMALL)
}

func Benchmark1MB(b *testing.B) {
	parseTestFile(b, "data/1MB.txt", 1, ONE_MB)
}

func Benchmark10MB1T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 1, TEN_MB)

}

func Benchmark10MB2T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 2, TEN_MB)

}

func Benchmark10MB4T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 4, TEN_MB)
}

func Benchmark10MB8T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 8, TEN_MB)
}

func Benchmark10MB16T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 16, TEN_MB)
}

func Benchmark10MB32T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 32, TEN_MB)
}
func Benchmark10MB64T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 64, TEN_MB)
}
func Benchmark10MB128T(b *testing.B) {
	parseTestFile(b, "data/10MB.txt", 128, TEN_MB)
}
