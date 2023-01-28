package main

import (
	"flag"
	"fmt"
)

func main() {
	patternsFile := flag.String("p", "", "The name of the pattern file")
	flag.Parse()

	//println(patternsFile)

	fmt.Println("Hello world ", *patternsFile)

}
