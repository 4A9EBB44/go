// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("os.Args[0] is:\n", os.Args[0])
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Println("val: ", i, " pointer: ", &i)
	}
}

//!-
