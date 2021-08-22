package main

// 입력과 출력이 들어있따
import "fmt"

// main 함수 : 프로그램이 실행될 때 호출되는 함수
func main() {
	fmt.Println("Hello World")

	var a int
	var f float32 = 11

	a = 10
	f = 12.0

	fmt.Println(a, f)

	// ''
	// :=는 변수 선언과 동시에 대입하는 연산자
	rawLiteral := `아리랑\n
아리랑\n
  아라리요`
	fmt.Println(rawLiteral)

	// ""
	interLiteral := "아리랑아리랑\n아리리요"
	fmt.Println(interLiteral)

	// 데이터 타입 변환
	var i int = 100
	var u uint = uint(i)
	var f_ float32 = float32(i)
	println(i, u, f_)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(str, bytes, str2)
}

