package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var LOGFILE1 = "/tmp/mGo1.log"
var LOGFILE2 = "/tmp/mGo2.log"
var PERMISSION = 0644

func main() {
	f1, err := os.OpenFile(LOGFILE1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	f2, err := os.OpenFile(LOGFILE2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()

	w := io.MultiWriter(f1, f2)
	iLog := log.New(w, " costumLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)

	iLog.Println("There he goes one of God's own prototypes, high-powered")
	iLog.Println("mutant of some king never even considered for mass prduction.")
	iLog.Println("too wierd to live, too wierd to die.")

}
