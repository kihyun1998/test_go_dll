package main

import (
	"fmt"
	"syscall"
)

// DLL에서 로드한 함수의 타입을 정의합니다
type addFunc = func(a, b int) int
type multiplyFunc = func(a, b int) int

func main() {
	// DLL 파일을 로드합니다
	// LoadLibrary와 비슷한 역할을 하는 LoadDLL을 사용합니다
	dll, err := syscall.LoadDLL("mathlib.dll")
	if err != nil {
		fmt.Printf("DLL 로드 실패: %v\n", err)
		return
	}
	defer dll.Release() // 프로그램 종료 전에 DLL을 해제합니다

	// Add 함수를 찾아서 프로시저로 가져옵니다
	addProc, err := dll.FindProc("Add")
	if err != nil {
		fmt.Printf("Add 함수를 찾을 수 없습니다: %v\n", err)
		return
	}

	// Multiply 함수를 찾아서 프로시저로 가져옵니다
	multiplyProc, err := dll.FindProc("Multiply")
	if err != nil {
		fmt.Printf("Multiply 함수를 찾을 수 없습니다: %v\n", err)
		return
	}

	// Add 함수 호출 테스트
	a, b := 5, 3
	result, _, _ := addProc.Call(
		uintptr(a),
		uintptr(b),
	)
	fmt.Printf("%d + %d = %d\n", a, b, result)

	// Multiply 함수 호출 테스트
	result, _, _ = multiplyProc.Call(
		uintptr(a),
		uintptr(b),
	)
	fmt.Printf("%d * %d = %d\n", a, b, result)
}
