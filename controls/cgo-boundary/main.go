package main

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

func main() {
	value := C.CString("opaque")
	defer C.free(unsafe.Pointer(value))
}
