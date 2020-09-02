package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sharedRPC"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		return
	}

	CONNECT := arguments[1]
	c, err := rpc.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	args := sharedRPC.MFloats{A1: 16, A2: -0.5}
	var reply float64

	err = c.Call("MInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
