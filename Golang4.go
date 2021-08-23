package main

func main() {
	// 익명 함수, Anonymous Function
	sum := func(n ...int) int {
		s := 0
		for _, i := range n{
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	println(result)


	// 일급함수
	// 함수는 기본 타입과 동일하게 취급
	// 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴 값으로 사용될 수 있다

	//변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int { return x + y }, 10, 20)
	println(r2)

	// type문
	r3 := calc_(func(x int, y int) int {return x + y}, 10, 20)
	println(r3)

	//클로저 사용
	next := nextValue()

	println(next())  // 1
	println(next())  // 2
	println(next())  // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2



}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

//type문을 사용한 함수 원형 정의

// 원형 정의
type calculator func(int, int) int

// calculator 원형 사용
func calc_(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

// 클로저 정의(Closure)
// 함수를 반환
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

