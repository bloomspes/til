package main

import (
	"fmt"
	"math/cmplx"
)

func personalData() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"peter", "luker", 37},
		{"kate", "young", 24},
		{"fred", "george", 18},
	}

	fmt.Println(people)
}

var (
	x    int
	y        = 20
	z    int = 30
	d, e     = 40, "hello"
	f, g string

	numbers = []int{10, 20, 30}
)

const (
	id       = "id"
	greeting = "반가워요, "
)

func sample() int {
	x = 10
	return x
}

func main() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)

	const prompt = "또 만나요!"

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	fmt.Println(greeting + id + "님 " + prompt)

	// 배열 끄트머리에 값 추가
	numbers = append(numbers, 50)
	// numbers := 20

	fmt.Println(numbers)

	result := make([]int, 4)
	result = append(result, 10)

	fmt.Println(result)

	initedSlice := make([]int, 5, 20)

	for i := 0; i < len(initedSlice); i++ {
		initedSlice[i] = 20
	}

	fmt.Println(initedSlice)

	fmt.Println(sample())

	fmt.Println(d)
}
