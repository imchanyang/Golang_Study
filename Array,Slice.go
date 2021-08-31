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

    // 슬라이스에 값 추가하기
    s = append(s, 4, 5, 6)
    fmt.Println(s)
    fmt.Println(len(s))
    fmt.Println(cap(s))

    // 슬라이스에 슬라이스를 붙일 때
    s = append(s, s_...)
    fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

    // 배열 예시
    S := [3]int{1, 2, 3}
    var c [3]int

    c = S
    c[0] = 5
    fmt.Println(S)
    fmt.Println(c)

    // 슬라이스 예시 reference
    S_ := []int {1, 2, 3}
    var c_ []int

    c_ = S_
    c_[0] = 5
    fmt.Println(S_)
    fmt.Println(c_)

    // 슬라이스 복사하기
    // 길이만큼 복사됨
    // 복사이기 때문에 t_를 바꿔도 t는 바뀌지 않음
    t := []int{1, 2, 3, 4, 5}
    t_ := make([]int, 3)

    copy(t_, t)
    fmt.Println(t)
    fmt.Println(t_)
    t_[0] = 7
    fmt.Println(t)
    fmt.Println(t_)


    fmt.Println(t[0:3])
    T := t[0:3]
    fmt.Println(T)
    T[0] = 100

    // 부분으로 슬라이스를 참조
    // 따라서 값이 같이 바뀐다
    fmt.Println(t)
    fmt.Println(T)


}