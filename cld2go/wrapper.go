package cld2go

// #include <stdlib.h>
// #include "wapper.h"
import "C"
import "unsafe"

// Detect detects the language by calling the underlying lib.
func Detect(text string) string {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))

	lang := C.DetectLang(cs, -1)

	var lang string
	if res != nil {
		lang = C.GoString(res)
	}
	return lang
}
