package main

import (
	"flag"
	"fmt"

	"github.com/wingyplus/go_headver"
)

var (
	headFlag = flag.Int("head", 0, "Number of head version")
	buildFlag = flag.Int("build", 0, "Number of build version")
	suffixFlag = flag.String("suffix", "", "Suffix after version")
)

func main() {
	flag.Parse()

	fmt.Println(headver.Build(*headFlag, *buildFlag, *suffixFlag))
}
