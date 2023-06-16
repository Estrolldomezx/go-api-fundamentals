package main

import (
	get "training/api"
	"training/apple"

	"fmt"
)

// variable
var a, b int
var c string
var e, f = 1, "e"

// struct
type Person struct {
	id   int
	age  int
	name string
	male string
}

func main() {

	var d bool
	var g = 20
	var h = float64(g) //float32
	var i = uint(h)

	//function
	apple.Apple()
	fmt.Println()
	fmt.Println(add(2, 3))
	fmt.Println(del(20, 3))
	fmt.Println(multipleAndDivide(5, 35))
	fmt.Println(mod(5, 10))
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, h, i)
	grade(55)
	isGreater(30, 20)

	//for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i : ", i+1)
	}

	// var sum = 1
	// for sum < 100 {
	// 	sum += sum
	// }

	// for { } infinite loop

	// defer หน่วงการทำงานจนกว่า func ที่ครอบ(main) จะทำงานเสร็จ
	defer fmt.Println("world ")
	// now world will show last line in terminal
	fmt.Println("Hello")

	a := 0
	defer fmt.Println("before change : ", a)
	a = 1
	fmt.Println("after change : ", a)

	c := Person{1, 20, "Dome", "male"}
	fmt.Println("id : ", c.id)
	fmt.Println("age : ", c.age)
	fmt.Println("name : ", c.name)
	fmt.Println("male : ", c.male)
	//print all in 1 line
	fmt.Println(c.id, c.age, c.name, c.male)

	// array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	// range
	ar := []int{3, 5, 7, 9}
	for i, r := range ar {
		fmt.Println(i+1, r)
	}

	// map
	m := map[int]string{
		1: "one",
	}

	fmt.Println(m[1])

	get.Request()
}

func add(x int, y int) int {
	return x + y
}

func del(x, y int) int {
	return x - y
}

func multipleAndDivide(x, y int) (int, int) {
	return x * y, x / y
}

func mod(x, y int) (mod int, div int) {
	mod = x % y
	div = x / y
	return
}

func grade(g int) {
	if g > 80 && g < 100 {
		fmt.Println("Excellent")
	} else if g > 50 && g < 80 {
		fmt.Println("Good")
	} else {
		fmt.Println("Fail")
	}
}

func isGreater(a, b int) {
	if sub := a - b; sub > 0 {
		fmt.Println("Greater")
	} else {
		fmt.Println("Smaller")
	}
}
