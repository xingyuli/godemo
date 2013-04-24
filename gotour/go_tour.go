/*
 * NOTE:
 * Inorder to make this file more clearer and no variable/function
 * conflictions, I declare variables and functions inside each
 * test as many as possible.
 */
package main

/*
 * This code groups the imports into a parenthesized, "factored" import
 * statement. You can also write multiple import statements, like:
 *     import "fmt"
 *     import "math"
 */
import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"math"
	"math/cmplx"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
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

// A function can take zero or more arguments.
func page6_Functions() {
	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(42, 13))
}

/*
 * When two or more consecutive named function parameters share a type,
 * you can omit the type from all but the last.
 */
func page7_FunctionsContinued() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(42, 13))
}

// A function can return any number of results.
func page8_MultipleResults() {
	swap := func(x, y string) (string, string) {
		return y, x
	}

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*
 * In Go, functions can return multiple "result parameters", not just a single
 * value. They can be named and act just like variables.
 *
 * If the result parameters are named, a return statement without arguments
 * returns the current values of the results.
 */
func page9_NamedResults() {
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}

	fmt.Println(split(17))
}

// The var statement declares a list of variables.
func page10_Variables() {
	var x, y, z int
	var c, python, java bool

	fmt.Println(x, y, z, c, python, java)
}

/*
 * A var declaration can include initializers, one per variable.
 *
 * If an initializer is present, the type can be omitted; the variable will
 * take the type of the initializer.
 */
func page11_VariablesWithInitializers() {
	var x, y, z int = 1, 2, 3
	var c, python, java = true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}

/*
 * Inside a function, the := short assignment statement can be used in place
 * of a var declaration with implicit type.
 *
 * (Outside a function, every construct begins with a keyword and the :=
 * construct is not available.)
 */
func page12_ShortVariableDeclarations() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}

/*
Go's basic types are:

bool

string

int int8 int16 int 32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32

float32 float64

complex64 complex128
*/
func page13_BasicTypes() {
	// NOTE: grouping can be used in functions as well!
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		comp   complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, comp, comp)
}

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
func page14_Constants() {
	const (
		World = "世界"
		Pi    = 3.14
		Truth = true
	)

	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)
}

/*
 * Numeric constants are high-precision values.
 *
 * An untyped constant takes the type needed by its context.
 */
func page15_NumericConstants() {
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	needInt := func(x int) int {
		return x*10 + 1
	}
	needFloat := func(x float64) float64 {
		return x * 0.1
	}

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// Big overflows int
	// fmt.Println(needInt(Big))
}

// Go has only looping construct, the for loop.
func page16_For() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// You can leave the pre and post statements empty.
func page17_ForContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func page18_ForIsGosWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func page19_Forever() {
	for {
	}
}

func page20_If() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
 * Question:
 * Is it possible to call the function itself like javascript's
 * "arguments.callee.call(args)" form?
 */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
 * Like for, the if statement can start with a short statement to execute
 * before the condition.
 *
 * Variabels declared by the statement are only in scope until the end of the
 * if.
 */
func page21_IfWithAShortStatement() {
	pow := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}

		// undefined: v
		// return v

		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
 * Variables declared inside an `if`'s short statement are also avaiable
 * inside any of the else blocks.
 */
func page22_IfAndElse() {
	pow := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("%g >= %g\n", v, lim)
		}
		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func page23_LoopsAndFunctions() {
	Sqrt := func(x float64) float64 {
		const delta = 1e-10

		z, previousZ := 0.5*x, float64(0)
		for math.Abs(z-previousZ) > delta {
			previousZ = z
			z = z - (z*z-x)/(2*z)
		}
		return z
	}

	fmt.Printf("Sqrt(2) = %v, math.Sqrt(2) = %v\n", Sqrt(2), math.Sqrt(2))
	fmt.Printf("Sqrt(3) = %v, math.Sqrt(3) = %v\n", Sqrt(3), math.Sqrt(3))
	fmt.Printf("Sqrt(4) = %v, math.Sqrt(4) = %v\n", Sqrt(4), math.Sqrt(4))
	fmt.Printf("Sqrt(5) = %v, math.Sqrt(5) = %v\n", Sqrt(5), math.Sqrt(5))
}

// A struct is a collectionof fields.
func page24_Structs() {
	fmt.Println(Vertex{1, 2})
}

type Vertex struct {
	X int
	Y int
}

// Struct fields are accessed using a dot.
func page25_StructFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

/*
 * Go has pointers, but no pointer arithmethic.
 *
 * Struct fields can be accessed through a struct pointer. The indirection
 * through the pointer is transparent.
 */
func page26_Pointers() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
	fmt.Printf("%T %v\n", p, p)
	fmt.Printf("%T %v\n", q, q) // q is a pointer to main.Vertex
}

/*
 * A struct literal denotes a newly allocated struct value by listing the
 * values of its fields.
 *
 * You can list just a subset of fields by using the Name: syntax.
 *
 * The special prefix & constructs a pointer to a struct literal.
 */
func page27_StructLiterals() {
	var (
		p = Vertex{1, 2}  // has type Vertex
		q = &Vertex{1, 2} // has type *Vertex
		r = Vertex{X: 1}  // Y:0 is implicit
		s = Vertex{}      // X:0 and Y:0
	)

	fmt.Println(p, q, r, s)
}

// The expression new(T) allocates a zeroed T value and returns a pointer to it.
func page28_TheNewFunction() {
	v := new(Vertex) // same with: var v *Vertex = new(Vertex)
	fmt.Println(v)

	v.X, v.Y = 11, 9
	fmt.Println(v)
}

/*
 * A slice points to an array of values and also includes a length.
 *
 * []T is a slice with elements of type T.
 */
func page29_Slices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

/*
 * Slices can be re-sliced, creating a new slice value that points to the
 * **same array**.
 */
func page30_SlicingSlices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	oneToFour := p[1:4]
	fmt.Println("\noneToFour := p[1:4]")
	fmt.Println("oneToFour ==", oneToFour) // [3, 5, 7]
	fmt.Println("p ==", p)                 // [2, 3, 5, 7, 11, 13]

	p[3] = -1
	/*
	 * In fact, oneToFour and p points to the same array. So any changes made
	 * to either slice is visible for other slices.
	 */
	fmt.Println("\np[3] = -1")
	fmt.Println("p ==", p)                  // [2, 3, 5, -1, 11, 13]
	fmt.Println("oneToFour == ", oneToFour) // [3, 5, -1]

	fmt.Println("\np[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])
}

/*
 * Slices are created with the make function. It works by allocating a zeroed
 * array and returning a slice that refers to that array.
 */
func page31_MakingSlices() {
	printSlice := func(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	}

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[0:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

/*
 * The zero value of a slice is nil.
 *
 * A nil slice has a length and capacity of 0.
 */
func page32_NilSlices() {
	var z []int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nil!")
	}
}

// The range form of the for loop iterates over a slice or map.
func page33_Range() {
	pow33 := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow33 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
 * You can skip the index or value by assigning to _.
 *
 * If you only want the index, drop the ", value" entirely.
 */
func page34_RangeContinued() {
	pow34 := make([]int, 10)
	for i := range pow34 {
		pow34[i] = 1 << uint(i)
	}
	for _, value := range pow34 {
		fmt.Printf("%d\n", value)
	}
}

/*
 * A map maps keys to values.
 *
 * Map must be created with make(not new) before use; the nil map is
 * empty and cannot be assigned to.
 */
func page36_Maps() {
	var m map[string]Vertex36
	fmt.Println("var m map[string]Vertex36\nm == nil?", m == nil) // true

	m = make(map[string]Vertex36)
	fmt.Println("\nm = make(map[string]Vertex36)\nm == nil?", m == nil) // false

	m["Bell Labs"] = Vertex36{
		40.68433, -74.39967,
	}
	fmt.Printf("\n%v", m["Bell Labs"])
}

type Vertex36 struct {
	Lat, Long float64
}

// Map literals are like struct literals, but the keys are required.
func page37_MapLiterals() {
	m := map[string]Vertex36{
		"Bell Labs": Vertex36{
			40.68433, -74.39967,
		},
		"Google": Vertex36{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

/*
 * If the top-level type is just a type name, you can omit it from the
 * elements of the literal.
 */
func page38_MapLiteralsContinued() {
	m := map[string]Vertex36{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func page39_MutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	/*
	 * If the key is not present, the result is the zero value for the map's
	 * element type.
	 */
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func page40_Exercise_Maps() {
	WordCount := func(s string) map[string]int {
		dict := make(map[string]int)
		for _, word := range strings.Fields(s) {
			dict[word] = dict[word] + 1
		}
		return dict
	}

	wc.Test(WordCount)
}

// Function are values too.
func page41_FunctionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

/*
 * And functions are **full closures**.
 *
 * The adder function returns a closure. Each closure is bound to its own sum
 * variable.
 */
func page42_FunctionClosures() {
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func page43_Exercise_FibonacciClosure() {
	// fibonacci is a function that returns a function that returns an int.
	fibonacci := func() func() int {
		a, b := 0, 1
		return func() int {
			a, b = b, a+b
			return b
		}
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
 * A case body breaks automatically, unless it ends with a fallthrough
 * statement.
 */
func page44_Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	testFallThrough := func(n int) {
		switch {
		case n >= 100:
			fmt.Println(">= 100")
			fallthrough
		case n >= 50:
			fmt.Println(">= 50")
			fallthrough
		case n >= 30:
			fmt.Println(">= 10")
			fallthrough
		case n >= 0:
			fmt.Println(">= 0")
		default:
			fmt.Println("Passed in:", n)
		}
	}
	testFallThrough(100)
	testFallThrough(-1)
}

/*
 * Switch cases evaluate cases from top to bottom, stopping when a case
 * succeeds.
 */
func page45_SwitchEvaluationOrder() {
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
}

/*
 * Switch without a condition is the same as switch true.
 *
 * This construct can be a clean way to write long if-then-else chains.
 */
func page46_SwitchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func page47_AdvancedExercise_ComplexCubeRoots() {
	Cbrt := func(x float64) float64 {
		z, previousZ, delta := x, float64(0), 1e-20
		for math.Abs(z-previousZ) > delta {
			previousZ = z
			z = z - (z*z*z-x)/(3*z*z)
		}
		return z
	}

	fmt.Println("Cbrt(2) =", Cbrt(2))
	fmt.Println("math.Pow(2, 1/3) = ", math.Pow(2, 1.0/3))
}

/*
 * Go does not have classes. However, you can define method on struct types.
 *
 * The method recevier appears in its own argument list between the func
 * keyword and the method name.
 */
func page49_Methods() {
	v := &Vertex49{3, 4}
	fmt.Println(v.Abs())
}

type Vertex49 struct {
	X, Y float64
}

func (v *Vertex49) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
 * In fact, you can define a method on **any type** you define in your package,
 * not just structs.
 *
 * You cannot define a method on a type from another package, or on a basic
 * type.
 */
func page50_MethodContinued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// error: cannot define new methods on non-local type time.Time.
/*
func (t *time.Time) Greeting() {
	fmt.Println("I am time.Time!");
}
*/

// error: cannot define new methods on non-local type int
/*
func (integer int) Succ() int {
	return integer + 1
}
*/

/*
 * Methods can be associated with a named type or a pointer the a named
 * type.
 *
 * There are two reasons to use a pointer receiver.
 *
 * 1st. to avoid copying the value on each method call (more efficient
 * if the value type is a large struct).
 *
 * 2nd. so that the method can modify the value that its receiver points
 * to.
 */
func page51_MethodsWithPointerReceivers() {
	v := &Vertex51{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

type Vertex51 struct {
	X, Y float64
}

func (v *Vertex51) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex51) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
 * An interface type is defined by a set of methods.
 *
 * A value of interface type can hold any value that implements those
 * methods.
 */
func page52_Interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex51{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex51 implements Abser
	// a = v  // a Vertex51, does not

	fmt.Println(a.Abs())
}

type Abser interface {
	Abs() float64
}

/*
 * A type implements an interface by implementing the methods.
 *
 * There is no explicit declaration of intent.
 *
 * Implicit interfaces decouple implementation packages from the packages
 * that define the interfaces: neither depends on the other.
 *
 * It also encourages the definition of precise interfaces, because you
 * don't have to find every implementation and tag it with the new
 * interface name.
 */
func page53_InterfacesAreSatisfiedImplicitly() {
	var w Writer

	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

type Reader interface {
	Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

/*
 * An error is anything that can describe itself as an error string. The idea
 * is captured by the predefined, built-in interface type, error, with its
 * single method, Error, returning a string:
 *
 *     type error interface {
 *         Error() string
 *     }
 *
 * The fmt package's various print routines automatically know to call the
 * method when asked to print an error.
 */
func page54_Errors() {
	run := func() error {
		return &MyError{time.Now(), "it did't work"}
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func page55_Exercise_Errors() {
	Sqrt := func(f float64) (float64, error) {
		if f < 0 {
			return 0, ErrNegativeSqrt(f)
		}
		return math.Sqrt(f), nil
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

/*
 * Package http serves HTTP requests using any value that implements
 * http.Handler:
 *
 *     package http
 *
 *     type Handler interface {
 *         ServeHTTP(w ResponseWriter, r *Request)
 *     }
 */
func page56_WebServers() {
	// Approach 1:
	// http.ListenAndServe("localhost:4000", new(Hello))

	// Approach 2:
	http.Handle("/hello", new(Hello))
	http.ListenAndServe("localhost:4000", nil)

	// Approach 3:
	/*http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! (via http.HandleFunc)")
	})
	http.ListenAndServe("localhost:4000", nil)*/
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func page57_Exercise_HTTPHandlers() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(s))
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

/*
 * A goroutine is a lightweight thread managed by the Go runtime.
 *
 *     go f(x, y, z)
 *
 * starts a new goroutine running
 *
 *     f(x, y, z)
 *
 * The evaluation of f, x, y and z happens in the current goroutine and the
 * execution of f happens in the new goroutine.
 */
func page62_Goroutines() {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}

	go say("world")
	say("hello")
}

func main() {
	page62_Goroutines()
}
