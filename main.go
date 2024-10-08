//go: main

//single line comment

/* multi
line
comment
*/

package main

import (
	"fmt"
	// . "fmt" // alias where we dont need to use fmt. prefix
	m "math"
	"os"
)

func main() {
	var a uint = 100
	const b uint = 20
	value := a / b // short declaration operator which infer  the type , declare and assign the value
	var value_float float64 = 10.5
	fmt.Println("Hello World")
	fmt.Println(value)
	fmt.Println(m.Round(value_float))
	fmt.Printf("value: %d\n", value)
	fmt.Printf("value_float: %.3f\n", value_float)
	fmt.Printf("value_float: %T\n", value_float)

	var _, _ int = 10, 20 // dummy variable
	var _ string = `this is an
        string literal`

	var _ string = "this is a string"
	x, y := testCaseGo()
	var sum int = x + int(y)
	fmt.Print("x is: ", x)
	fmt.Printf("\t")
	fmt.Print("y is: ", y)
	fmt.Println("\tSum is: ", sum)
	byteValue := byte(65) // alais for uint8
	fmt.Println(byteValue)
	fmt.Printf("%T\n", byteValue)
	runeValue := rune(65) // alais for int32
	fmt.Printf("%T\n", runeValue)

	/* array */

	var _ [5]int = [5]int{1, 2, 3, 4, 5}

	var _ [5]int

	_ = [...]int{1, 2, 3, 4, 5}

	oneTofive := [...]int{1, 2, 3, 4, 5}

	fmt.Println(oneTofive)

	// slicing

	/*
		print and println are the lowaer level functions and fmt.Println is the higher level function which internally calls the print and println functions where println and print might not support some types like slice, map, struct etc. and might return references
	*/

	oneTofive_cpy := oneTofive // copy of the array
	oneTofive_cpy[0] = 100
	println(oneTofive[0] == oneTofive_cpy[0])

	_ = oneTofive[1:3] // slicing from index 1 to 3

	// slice are different than array as they are dynamic in nature

	s := []int{1, 2, 3, 4, 5} // slice
	s = append(s, 6)          // append element to the slice

	println("slice:", s, "length:", len(s), "capacity:", cap(s)) // here while printing the slice the reference value is returned
	fmt.Println("slice:", s)

	// pointers

	var p *uint
	p = &a
	fmt.Println(*p, p)

	q := &a
	fmt.Println(*q, q)

	pr, qr, rr := testMemomryGo()
	fmt.Println(*pr, *qr, *rr)
	println("pr:", pr, "qr:", qr, "rr:", rr)

	// maps

	/* maps can grow dynamically and are of type hash, dict, associative array . */

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	var m1 = map[string]int{"one": 1, "two": 2, "three": 3}
	println(m["one"], m1["two"])

	// alocating and initiating for map

	var m2 = make(map[string]int)
	m2["one"] = 1
	m2["two"] = 2

	m3 := make(map[string]int)
	m3["one"] = 1

	fmt.Println(m2, m3)
	println(m2, m3) //return memory address

	if value, ok := m3["one"]; ok {
		fmt.Println(value)
	}

	// file handeling

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(file, "Hello World")
	defer file.Close()
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	print(string(content))
	/*
		delay the execution of the close function until the end of the function which means the file will be closed at the end of the main.
		if there are multiple defer function follows the LIFO principle as the last defer function is called first before closing the end of the functions
	*/

	// flow control

	// if statement
	if false {
	} else if false {
	} else {
		println("else")
	}

	// switch statements

	val := 10.0

	switch int(val) {
	case 10:
		fmt.Println("value is 10")
		fallthrough // fallthrough to the next case its like after case satisfied eventhoug the next case is not satisfied it will execute the next case
	default:
		fmt.Println("value is not 10")
	}

	// type switch  case

	var type1 interface{}
	// the type is supported because the type1 is a interface type

	switch type1.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case fmt.Stringer:
		fmt.Println("type is fmt.Stringer interface")
	default:
		fmt.Println("type not found")
	}

	for x := 0; x < 10; x++ {
		print(x)
	}
	println()

	var i int = 10
	for ; i > 0; i-- {
		print(i)
	}
	println()

	var list [5]int = [...]int{1, 2, 3, 4, 5}
	for indx, v := range list { // range returns the index and value of the list
		fmt.Printf("index: %d, value: %d\n", indx, v)
	}

	for {
		fmt.Println("infinite loop")
		break
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key: %s, value: %d, type:%T\n", key, value, value)
	}

	xBigThanY := func() bool { //anonymous function
		return x > int(y)
	}
	println(xBigThanY())
	x = 1000
	println(xBigThanY())

	//label and goto

	var counter int = 0

label:
	println("label")
	counter++
	if counter < 2 {
		goto label
	}
	testFunctionDecorators()
	testStruct()

}

func testCaseGo() (x int, y uint) {
	x = 10
	y = 20
	return
}

func testMemomryGo() (*int, *int, *int) {
	var alloc *int = new(int) //allocate memory for type int and return a pointer to it
	slice := make([]int, 5)   //contiguous  allocate memory
	slice[2] = 10
	value := 23
	alloc = &value
	return alloc, &slice[1], &slice[2]
}

func testFunctionDecorators() {
	var input string
	fmt.Println("Enter the input string: ")
	fmt.Scanf("%s", &input)
	println(dec_func(input)("before", "after"))
	//equivalence of the above statement
	decorator := dec_func(input)
	println(decorator("before", "after"))
}

func dec_func(input string) func(before, after string) string {
	println("decorator function input is: ", input)
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, input, after)
	}
}

type TestInter interface {
	returnName() (string, string)
}

type TestStruct struct {
	name string
	age  int
}

// implementing the method for the struct
// without interface implementation the method is called as the method of the struct
func (t *TestStruct) testFunc() int { // pointer receiver as it works on the same instance of the struct
	return t.age
}

// interface implementation as the method is called as the method of the interface
func (t TestStruct) returnName() (actual_name string, edited_name string) { // value receiver as it works on the copy of the struct
	actual_name = t.name
	t.name = "edited"
	edited_name = t.name
	return
}
func testStruct() {
	var t1 TestStruct = TestStruct{name: "test", age: 10}
	println(t1.testFunc())
	n1, n2 := t1.returnName()
	println(n1, n2)
	println(t1.name)

	//reusable code with interface
	var t2 TestInter
	t2 = t1 // as t1 implements the TestInter interface
	println(t2.returnName())

        vardicParm(10, "string") // vardic parametres in a function
}

// vardic parametres in a function
func vardicParm(input ...any) {
	for _, value := range input {
		fmt.Println(value)
        }	
}






