package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("pass at least one number argument")
		return
	}

	args := os.Args
	var sum float64 = 0
	var err error = errors.New("error")
	i := 1
	for err != nil {
		if i >= len(args) {
			fmt.Println("pass at least one number argument")
			return
		}
		sum, err = strconv.ParseFloat(args[i], 64)
		i++
	}

	for i < len(args) {
		val, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			sum += val
		}
		i++
	}

	fmt.Println("sum = ", sum)
}
