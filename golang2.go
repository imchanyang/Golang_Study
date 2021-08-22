package main

func main() {

	//for문 : 기본적인 반복문
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	//for문 : 조건식만 쓰는 반복문, while이랑 비슷하다
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)

	//for문 : 무한루프
	//for {
	//	println("Infinite loop")
	//}

	//for문 : range를 이용한 반복문
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		println(index, name)
	}

	//for문 : break, continue, goto
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break  //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

	END:
	println("End")
}


