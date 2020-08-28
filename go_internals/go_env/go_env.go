package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("you are useing ", runtime.Compiler)
	fmt.Println("on a ", runtime.GOARCH, " machine")
	fmt.Println("useing go version ", runtime.Version())
	fmt.Println("Number of CPUs ", runtime.NumCPU())
	fmt.Println("number of gorutines ", runtime.NumGoroutine())
}
