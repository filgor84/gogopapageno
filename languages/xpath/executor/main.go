package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/filgor84/gogopapageno/languages/xpath"
)

const fileToParsePath = "./data/small.1.xml"

var (
	q = flag.String("q", "A1", "query name")
	f = flag.String("f", "", "file name")
	n = flag.Int("n", 1, "number of threads")
	v = flag.Bool("v", false, "verbose mode")
)

var usage = `Usage parserTester
	-q	Name of the query type to be executed
	-f	Name of the file containing the XML document to query
	-n	Number of parsers to run concurrently
	-v  In verbose mode
`

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()

	if *f == "" {
		usageAndExit("-f can not be empty")
	}

	command := xpath.Execute(*q).AgainstFile(*f).WithNumberOfThreads(*n)

	if *v {
		command = command.InVerboseMode()
	}

	results, err := command.Go()

	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot execute command: %v", err)
		fmt.Fprintf(os.Stderr, "\n\n")
		os.Exit(1)
	}

	fmt.Printf("results: %v\n", results)
	fmt.Printf("number of matches: %d\n", len(results))

}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
