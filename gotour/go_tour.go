package main

import (
  "fmt"
	"math"
	"math/cmplx"
)

func page1_HelloWorld() {
	fmt.Println("Hello World!")
}


func page3_Packages() {
	fmt.Println("Happy", math.Pi, "Day")
}


func page4_Imports() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}


func page5_ExportedNames() {
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)
}


func page6_Functions() {
	fmt.Println(add(42, 13))
}
func add(x int, y int) int {
	return x + y
}


func page7_FunctionsContinued() {
	fmt.Println(add(42, 13))
}
func add2(x, y int) int {
	return x + y
}


func page8_MultipleResults() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
func swap(x, y string) (string, string) {
	return y, x
}


func page9_NamedResults() {
	fmt.Println(split(17))
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


func page10_Variables() {
	fmt.Println(x, y, z, c, python, java)
}
var x, y, z int
var c, python, java bool


func page11_VariablesWithInitializers() {
	fmt.Println(x1, y1, z1, c1, python1, java1)
}
var x1, y1, z1 int = 1, 2, 3
var c1, python1, java1 = true, false, "no!"


func page12_ShortVariableDeclarations() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}


func page13_BasicTypes() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, comp, comp)
}
var (
	ToBe	bool		= false
	MaxInt	uint64		= 1<<64 - 1
	comp	complex128	= cmplx.Sqrt(-5 + 12i)
)


func page14_Constants() {
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
const Pi = 3.14


func page15_NumericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// Big overflows int
	// fmt.Println(needInt(Big))
}
const (
	Big		= 1 << 100
	Small	= Big >> 99
)
func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}


func page16_For() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}


func page17_ForContinued() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	page17_ForContinued()
}
