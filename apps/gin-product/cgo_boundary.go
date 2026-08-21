package main

/*
#include <stdlib.h>
#include <string.h>

static size_t sivere_copy_checked(char *destination, size_t capacity,
                               const char *source, size_t length) {
    if (length >= capacity) {
        return 0;
    }
    memcpy(destination, source, length);
    destination[length] = '\0';
    return length;
}

static size_t sivere_copy_unchecked(char *destination,
                                 const char *source, size_t length) {
    memcpy(destination, source, length);
    destination[length] = '\0';
    return length;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// cgoBoundaryOperation crosses a real CGO boundary in both controls. The safe
// control supplies and enforces the destination capacity. The vulnerable
// control calls an API whose contract has no capacity argument. Allocation is
// deliberately large enough for the qualification probe so the benchmark
// observes the unsafe contract without corrupting the test process.
func cgoBoundaryOperation(value string, protected bool) string {
	source := C.CString(value)
	defer C.free(unsafe.Pointer(source))
	destination := C.malloc(C.size_t(len(value) + 1))
	if destination == nil {
		return "allocation_failed"
	}
	defer C.free(destination)
	if protected {
		length := C.sivere_copy_checked(
			(*C.char)(destination),
			C.size_t(len(value)+1),
			source,
			C.size_t(len(value)),
		)
		return fmt.Sprintf("checked:%d", uint64(length))
	}
	length := C.sivere_copy_unchecked(
		(*C.char)(destination),
		source,
		C.size_t(len(value)),
	)
	return fmt.Sprintf("unchecked:%d", uint64(length))
}
