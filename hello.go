package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/jimpick/sp-kyc-go-test-simple/morestrings"
)

// Tutorial: https://go.dev/doc/code

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
