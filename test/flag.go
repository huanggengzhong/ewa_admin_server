package main

import (
	"flag"
	"fmt"
)

var (
	intFlag    int
	boolFlag   bool
	stringFlag string
)

func init() {
	flag.IntVar(&intFlag, "intFlag", 0, "int flag value")
	flag.BoolVar(&boolFlag, "boolFlag", false, "bool flag value")
	flag.StringVar(&stringFlag, "stringFlag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("int flag:", intFlag)
	fmt.Println("bool flag:", boolFlag)
	fmt.Println("string flag:", stringFlag)
}
