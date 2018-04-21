package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool
var i, j int = 1, 2

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func substract(x, y int) int {
	return x - y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problems", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	fmt.Print(substract(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(18))
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, c, python, java)
	var i int
	fmt.Println(i, k, c, python, java)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
