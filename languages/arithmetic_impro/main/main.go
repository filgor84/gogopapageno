package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/filgor84/gogopapageno/languages/arithmetic_impro"
)

var cpuprofile = flag.String("cpuprofile", "", "") //write cpu profile to file")
var memprofile = flag.String("memprofile", "", "") //write memory profile to file")

var cpuprofileFile *os.File

var fname = flag.String("fname", "", "the name of the file to parse")
var numThreads = flag.Int("n", 1, "the number of threads to use")

func main() {
	//Set flags (for debugging only)
	//flag.Set("fname", "languages/arithmetic_impro/data/small.txt")
	//flag.Set("n", "2")

	//Set the usage message that is printed when incorrect or insufficient arguments are passed
	flag.Usage = func() {
		fmt.Println("Usage: main -fname filename [-n numthreads]")
	}

	flag.Parse()

	if *fname == "" || *numThreads < 1 {
		flag.Usage()
		return
	}

	//Code needed for the cpu profiler
	if *cpuprofile != "" {
		err := error(nil)
		cpuprofileFile, err = os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		arithmetic_impro.SetCPUProfileFile(cpuprofileFile)
	}

	fmt.Println("Available cores:", runtime.GOMAXPROCS(0))

	fmt.Println("Number of threads:", *numThreads)
	start := time.Now()

	root, err := arithmetic_impro.ParseFile(*fname, *numThreads)
	total_time := time.Since(start)

	if err == nil {
		fmt.Println("Parse succeded!")
		for i, v := range arithmetic_impro.Stats.StackPoolSizes {
			fmt.Printf("Stack pool size (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic_impro.Stats.StackPoolNewNonterminalsSizes {
			fmt.Printf("Stack pool new nonterminals size (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic_impro.Stats.StackPtrPoolSizes {
			fmt.Printf("StackPtr pool size (thread %d): %d\n", i, v)
		}
		fmt.Printf("Stack pool final pass size: %d\n", arithmetic_impro.Stats.StackPoolSizeFinalPass)
		fmt.Printf("Stack pool final pass new nonterminals size: %d\n", arithmetic_impro.Stats.StackPoolNewNonterminalsSizeFinalPass)
		fmt.Printf("StackPtr pool final pass size: %d\n", arithmetic_impro.Stats.StackPtrPoolSizeFinalPass)
		fmt.Printf("Time to alloc memory: %s\n\n", arithmetic_impro.Stats.AllocMemTime)

		for i, v := range arithmetic_impro.Stats.CutPoints {
			fmt.Printf("cutpoint %d: %d\n", i, v)
		}
		for i, v := range arithmetic_impro.Stats.LexTimes {
			fmt.Printf("Time to lex (thread %d): %s\n", i, v)
		}
		fmt.Printf("Time to lex (total): %s\n\n", arithmetic_impro.Stats.LexTimeTotal)

		for i, v := range arithmetic_impro.Stats.NumTokens {
			fmt.Printf("Number of tokens (thread %d): %d\n", i, v)
		}
		fmt.Printf("Number of tokens (total): %d\n", arithmetic_impro.Stats.NumTokensTotal)
		for i, v := range arithmetic_impro.Stats.ParseTimes {
			fmt.Printf("Time to parse (thread %d): %s\n", i, v)
		}
		fmt.Printf("Time to recombine the stacks: %s\n", arithmetic_impro.Stats.RecombiningStacksTime)
		fmt.Printf("Time to parse (final pass): %s\n", arithmetic_impro.Stats.ParseTimeFinalPass)
		fmt.Printf("Time to parse (total): %s\n\n", arithmetic_impro.Stats.ParseTimeTotal)

		for i, v := range arithmetic_impro.Stats.RemainingStacks {
			fmt.Printf("Remaining stacks (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic_impro.Stats.RemainingStacksNewNonterminals {
			fmt.Printf("Remaining stacks new nonterminals (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic_impro.Stats.RemainingStackPtrs {
			fmt.Printf("Remaining stackPtrs (thread %d): %d\n", i, v)
		}

		fmt.Printf("Remaining stacks final pass: %d\n", arithmetic_impro.Stats.RemainingStacksFinalPass)
		fmt.Printf("Remaining stacks new nonterminals final pass: %d\n", arithmetic_impro.Stats.RemainingStacksNewNonterminalsFinalPass)
		fmt.Printf("Remaining stackPtrs final pass: %d\n\n", arithmetic_impro.Stats.RemainingStackPtrsFinalPass)

		fmt.Printf("Result: %d\n", *root.Value.(*int64))
		fmt.Printf("Total execution time: %f secs\n", float64(total_time)/1000000000.0)
	} else {
		fmt.Println("Parse failed!")
		fmt.Println(err.Error())
	}

	//Code needed for the mem profiler
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		//runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
