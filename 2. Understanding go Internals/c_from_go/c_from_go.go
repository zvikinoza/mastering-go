package main

// #cgo CFLAGS: -I${SRCDIR}/Clib
// #cgo LDFLAGS: ${SRCDIR}/rinter.a
// #include <stdlib.h>
// #include <rinter.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Printing from main go")
	C.rintC()

	fmt.Println("begin call C.printmessage")
	msg := C.CString("sic mundus creatus est")
	defer C.free(unsafe.Pointer(msg))
	C.rintCMsg(msg)

	fmt.Println("end call C.printmessage")
}
