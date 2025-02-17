package main

/*
#include <stdlib.h>
typedef struct {
    char *msg;
    int isOk;
} GoResult;
*/
import "C"
import (
	"fmt"
	"strings"
)

//export Add
func Add(a, b int) int {
	return a + b
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

//export SayHello
func SayHello(name *C.char) *C.char {
	goName := C.GoString(name)
	result := fmt.Sprintf("Hello, %s!", goName)
	return C.CString(result)
}

//export CheckBook
func CheckBook(result *C.GoResult, bookName *C.char) {
	goBookName := C.GoString(bookName)
	var msg string
	var ok int

	if strings.Contains(strings.ToLower(goBookName), "golang") {
		msg = fmt.Sprintf("Book '%s' is available", goBookName)
		ok = 1
	} else {
		msg = fmt.Sprintf("Sorry, '%s' is not available", goBookName)
		ok = 0
	}

	result.msg = C.CString(msg)
	result.isOk = C.int(ok)
}

func main() {}
