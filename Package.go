package main

// 이렇게 나열할 수 있다.
import (
	// 전역 패키지 fmt를 생략하고 사용할 수 있다.
	. "fmt"
	"io/ioutil"

	// 패키지 별칭 사용하기
	 a "runtime"
)

func main() {
	// fmt 생략
	Println("CPU count : ", a.NumCPU())

	var b[] byte
	var err error

	b, err = ioutil.ReadFile("./hello.txt")

	if err == nil {
		Printf("%s", b)
	}
}
