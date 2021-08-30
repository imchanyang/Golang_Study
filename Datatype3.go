package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 10.0

	for i := 0; i <10; i++ {
		a = a - 0.1
	}

	// 2진수로는 실수를 정확하게 표현할 수 없다. 따라서 오차가 발생한다.
	// 9.0000000000004정도의 값이 나와서
	// 실수는 연산한 값을 비교하면 잘못된 결과가 나오므로 주의해야한다.
	fmt.Println(a)
	fmt.Println(a == 9.0)

	// 실수를 비교하려면 연산한 값과 비교할 값의 차이를 구한 뒤 머신 입실론 보다 작거나 같은지 확인하면 된다.
	const epsilon = 1e-14
	fmt.Println(math.Abs(a-9.0) <= epsilon)

	// 복소수 실수부 + 허수부
	var num1 complex64 = 1 + 2i
	fmt.Println(real(num1))
	fmt.Println(imag(num1))

	var num2 complex64 = complex(1, 2)
	fmt.Println(real(num2))
	fmt.Println(imag(num2))

	// 바이트
	var num_ byte = 10
	fmt.Println(num_)

	//룬 유니코드 (UTF-8) 문자 코드를 저장할 때 사용
	var r1 rune = '한'
	var r2 rune = '\ud55c'
	var r3 rune = '\U0000d55c'

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)

	// Go에서는 오버플로우, 언더플로우를 허용하지 않는다
	// var num3 uint8 = math.MaxUint8 + 1
	// var num3 uint8 = 0 - 1
}