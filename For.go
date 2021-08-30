package main

import "fmt"

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


	// Loop:
	// Loop:와 for문 사이에 무언가가 있으면 에
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")


Loop2:
	for i := 0; i < 3; i++ {              // 반복문 1
		for j := 0; j < 3; j++ {      // 반복문 2
			if j == 2 {           // j가 2일 때
				continue Loop2 // 아래 부분 코드를 실행하지 않고 반복문 1부터 이어서 실행
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")


	// for문에서 변수 여러 개 사용
	for i, j := 0, 0; i < 10; i, j = i + 1, j + 2 {
		fmt.Println(i, j)
	}
}


