package main

func main() {

	println(add(2, 3))

	msg := "Hello"
	say(&msg)
	println(msg)

	say_("This", "is", "a", "book")
	say_("Hi")

	count, total := sum(1, 7, 3, 5, 9)
	println(count, total)

	count_, total_ := sum_(1, 7, 3, 5, 9)
	println(count_, total_)
}

// 함수 변수타입, 반환 타입 쓰는 법
// Pass By Value
func add(a int, b int) int {
	return a + b
}

// Pass By Reference
func say(msg *string) {
	println(*msg)
	*msg = "Changed"
}

// 가변 인자 함수
// ...(3개의 마침표를 사용) 0개 ~ n개의 동일 타입 파라미터를 전달할 수 있다
func say_(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

// 복수 개의 값을 리턴하는 함수
func sum(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

// Named Return Parameter
func sum_(nums ...int) (count int, total int) {
	for _,n := range nums {
		total += n
	}
	count = len(nums)
	return
}