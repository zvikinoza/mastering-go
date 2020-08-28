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
	var sum float64
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

	cnt := 1
	for i < len(args) {
		val, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			sum += val
			cnt++
		}
		i++
	}

	fmt.Println("sum = ", sum/float64(cnt))
}
