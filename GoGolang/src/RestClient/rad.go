package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool
var i, j int = 1, 2
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func add(x int, y int) int {
	return x + y
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
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

type Vertex struct {
	X int
	Y int
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a1 := names[0:2]
	b1 := names[1:3]
	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)
	fmt.Println(a1, b1)
	fmt.Println(v1, p, v2, v3)
	var myarray [2]string

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[1:4]
	fmt.Println(s)
	myarray[0] = "hello"
	myarray[1] = "world"
	fmt.Println(myarray[0], myarray[1])
	fmt.Println(myarray)
	ii, jj := 42, 2701
	l := &ii
	fmt.Println(*l)
	*l = 21
	fmt.Println(ii)
	l = &jj
	*l = *l / 37
	fmt.Println(jj)
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	m := &v
	m.X = 1e9
	fmt.Println(v)
	fmt.Println(v.X)
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
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	defer fmt.Println("world")

	fmt.Println("hello")

}
