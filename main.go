package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// C의 GoResult (typedef struct { char *msg; int isOk; })는 64비트에서
// 포인터(8바이트) + int(4바이트) + 4바이트 패딩 = 16바이트입니다.
// 따라서 패딩을 추가하여 메모리 레이아웃을 맞춥니다.
type GoResult struct {
	Msg  uintptr
	IsOk int32
	_    [4]byte // 패딩: 총 크기를 16바이트로 맞춤
}

// 널 종료 C 문자열을 Go 문자열로 변환하는 함수
func cStringToGoString(cString uintptr) string {
	var result []byte
	ptr := (*byte)(unsafe.Pointer(cString))
	for *ptr != 0 {
		result = append(result, *ptr)
		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
	}
	return string(result)
}

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
	addResult, _, _ := addProc.Call(uintptr(a), uintptr(b))
	fmt.Printf("%d + %d = %d\n", a, b, addResult)

	// SayHello 함수 테스트
	sayHelloProc, err := dll.FindProc("SayHello")
	if err != nil {
		fmt.Printf("SayHello 함수를 찾을 수 없습니다: %v\n", err)
		return
	}
	name := "Gopher"
	namePtr, err := syscall.BytePtrFromString(name)
	if err != nil {
		fmt.Printf("문자열 변환 실패: %v\n", err)
		return
	}
	helloPtr, _, _ := sayHelloProc.Call(uintptr(unsafe.Pointer(namePtr)))
	helloStr := cStringToGoString(helloPtr)
	fmt.Println(helloStr)

	// CheckBook 함수 테스트 (출력 구조체 전달 방식)
	checkBookProc, err := dll.FindProc("CheckBook")
	if err != nil {
		fmt.Printf("CheckBook 함수를 찾을 수 없습니다: %v\n", err)
		return
	}
	bookName := "Golang Programming"
	bookPtr, err := syscall.BytePtrFromString(bookName)
	if err != nil {
		fmt.Printf("문자열 변환 실패: %v\n", err)
		return
	}
	// 미리 결과를 담을 메모리(구조체)를 준비합니다.
	var resultBook GoResult

	// CheckBook의 첫 번째 인자로 결과 구조체의 주소, 두 번째 인자로 책 이름 문자열 포인터를 전달합니다.
	checkBookProc.Call(uintptr(unsafe.Pointer(&resultBook)), uintptr(unsafe.Pointer(bookPtr)))

	resultMsg := cStringToGoString(resultBook.Msg)
	// isOk는 0 또는 1이므로 bool로 변환
	isOk := (resultBook.IsOk != 0)
	fmt.Printf("[Book] isOK: %t\n", isOk)
	fmt.Printf("[Book] msg: %s\n", resultMsg)
}
