package main

func main() {
	// 포인터 연산자
	var k int = 10
	// k의 주소 할당
	var p = &k
	// p가 가리키는 주소에 있는 실제 내용을 출력
	println(*p)

	//if문 : 조건문
	// else if 나 else는 형식을 맞춰야한다
	if k == 1{
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("other")
	}

	// 조건식을 넣기전에 Optional Statement를 함께 실행할 수 있다.
	i := 1
	max := 5
	if val := i * 2; val < max {
		println(val)
	}

	//Switch문 : 조건문
	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
		// break를 하지 않고 밑의 case문을 실행한다
		fallthrough
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	// 변수의 타입을 검사하는 방법
	var x interface{} = "foo"
	switch x.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}

}
