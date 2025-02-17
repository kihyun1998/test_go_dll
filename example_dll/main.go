// example/main.go
package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unicode/utf8"
)

//export SayHello
func SayHello(name *C.char) (*C.char, bool) {
	goName := C.GoString(name)
	if !utf8.ValidString(goName) {
		return C.CString("Error: Invalid UTF-8 string"), false
	}
	if goName == "" {
		return C.CString("Error: Name is empty"), false
	}

	result := fmt.Sprintf("안녕하세요, %s님!", goName)
	cstr := C.CString(result)
	return cstr, true
}

//export RepeatString
func RepeatString(text *C.char, count C.int) (*C.char, bool) {
	goText := C.GoString(text)
	if !utf8.ValidString(goText) {
		return C.CString("Error: Invalid UTF-8 string"), false
	}
	if goText == "" {
		return C.CString("Error: Text is empty"), false
	}

	if count <= 0 {
		return C.CString("Error: Count must be positive"), false
	}

	result := ""
	for i := 0; i < int(count); i++ {
		result += goText
	}

	cstr := C.CString(result)
	return cstr, true
}

func main() {}
