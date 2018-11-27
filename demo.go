package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

//var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {

	//method2
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	//cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {

	//method 1
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}

	//method 3
	//cmdLine.Parse(os.Args[1:])
	flag.Parse()
	fmt.Printf("Hello %s!\n", name)
}
