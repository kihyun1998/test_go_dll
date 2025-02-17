package main

import "C"
import "fmt"

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

func main() {
}
