package main

import "fmt"
// import로 패키지를 가져올 떄도 쓰지 않으면 에러가 발생한다.
// 패키지 이름 앞에 _를 붙이면 에러가 발생하지 않는다.
import _ "math"

func main() {

	// Go언어는 선언만 하고 사용하지 않는 변수가 있으면 에러가 발생한다.
	// 어쩔 수 없이 사용하지 않는 변수를 만들때는 _로 변수명을 사용하면 에러가 발생하지 않는다.
	a := 1
	b := 2
	_ = b

	fmt.Println(a)
}
