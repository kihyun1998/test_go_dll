package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	// DLL 로드
	dll, err := syscall.LoadDLL("mathlib.dll")
	if err != nil {
		fmt.Printf("DLL 로드 실패: %v\n", err)
		return
	}
	defer dll.Release()

	// Add 함수 테스트
	addProc, err := dll.FindProc("Add")
	if err != nil {
		fmt.Printf("Add 함수를 찾을 수 없습니다: %v\n", err)
		return
	}

	a, b := 5, 3
	result, _, _ := addProc.Call(
		uintptr(a),
		uintptr(b),
	)
	fmt.Printf("%d + %d = %d\n", a, b, result)

	// SayHello 함수 테스트
	sayHelloProc, err := dll.FindProc("SayHello")
	if err != nil {
		fmt.Printf("SayHello 함수를 찾을 수 없습니다: %v\n", err)
		return
	}

	// 테스트할 문자열
	name := "Gopher"
	// syscall.StringBytePtr를 사용하여 문자열을 C 스타일 바이트 포인터로 변환
	namePtr, err := syscall.BytePtrFromString(name)
	if err != nil {
		fmt.Printf("문자열 변환 실패: %v\n", err)
		return
	}

	// SayHello 함수 호출
	resultPtr, _, _ := sayHelloProc.Call(uintptr(unsafe.Pointer(namePtr)))

	// C 문자열을 Go 문자열로 변환
	// C 문자열은 null로 끝나므로, null을 만날 때까지 읽어서 변환합니다
	result_str := ""
	ptr := (*byte)(unsafe.Pointer(resultPtr))
	for *ptr != 0 {
		result_str += string(*ptr)
		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
	}

	fmt.Println(result_str)
}
