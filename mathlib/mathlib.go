package main

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
	// Go 문자열로 변환
	goName := C.GoString(name)
	// 새로운 문자열 생성
	result := fmt.Sprintf("Hello, %s!", goName)
	// C 문자열로 변환하여 반환
	return C.CString(result)
}

//export CheckBook
func CheckBook(bookName *C.char) (msg *C.char, isOk bool) {
	goBookName := C.GoString(bookName)

	// "golang"이 포함된 경우 책이 있다고 가정
	if contains := strings.Contains(strings.ToLower(goBookName), "golang"); contains {
		return C.CString(fmt.Sprintf("Book '%s' is available", goBookName)), true
	}

	return C.CString(fmt.Sprintf("Sorry, '%s' is not available", goBookName)), false
}

func main() {
}
