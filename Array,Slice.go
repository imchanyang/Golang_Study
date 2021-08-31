package main

import "fmt"

func main() {
	var a [5]int
	a[2] = 5

	fmt.Println(a)

	var b = [3]int{1, 2, 3}
	fmt.Println(b)

	d := [...]string{"as", "sdf", "asdf"}
	fmt.Println(d)

	// 다차원 배열
	a_ := [2][2]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(a_)

	b_ := [5]int{1, 2, 3, 4, 5}
	for i, value := range (b_) {
		fmt.Println(i, value)
	}

	// 슬라이스
	// 배열과 비슷하지만 길이가 고정되어 있지 않고 동적으로 크기가 늘어난다.
	// reference 타입

	// make 함수로 int형에 길이가 5인 슬라이스에 공간 할당
	var s []int = make([]int, 5)
	// 같은 표현들
	// var s = make([]int, 5)
	// s := make([]int, 5)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	// 슬라이스를 생성하면서 값을 초기화
	s_ := []int{1, 2, 3, 4, 5}
	fmt.Println(s_)
	fmt.Println(len(s_))
	fmt.Println(cap(s_))

	// 용량이 10인 슬라이스
	// 용량이 가득차면 용량은 자동으로 늘어난다
	// 인덱스로 접근할 때 길이를 벗어나면 에러
    var s__ = make([]int, 5, 10)
    fmt.Println(len(s__))
    fmt.Println(cap(s__))

    s = []int {1, 2}
    fmt.Println(s)
    fmt.Println(len(s))
    fmt.Println(cap(s))

}