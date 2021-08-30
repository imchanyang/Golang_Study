package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var num1 int8  = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	//바이트 단위의 변수의 크기가 나온다
	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))


	// 문자열 길이 구하기
	var s1 string = "한글"
	var s2 string = "Hello"
	// 한글을 UTF-8로 저장하면 길이가 6이다
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	// 실제 크기를 구하려면
	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Printf("%c\n", s2[1])

	// 상수 열거형 사용하기
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Satuday
		numberOfDays
	)
	fmt.Println("-----------")
	fmt.Println(Tuesday)
}
