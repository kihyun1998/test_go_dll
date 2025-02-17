package main

import "C"

//export Add
func Add(a, b int) int {
	return a + b
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
	// DLL에는 main 함수가 필요하지만 비워둡니다
}
