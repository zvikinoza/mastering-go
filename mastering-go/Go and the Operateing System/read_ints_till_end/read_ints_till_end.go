package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "END" {
			return
		}
		fmt.Println(txt)
	}
}
