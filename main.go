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
	fmt.Printf("%T", byteValue)
	runeValue := rune(65) // alais for int32
	fmt.Printf("%T", runeValue)

	/* array */

	var _ [5]int = [5]int{1, 2, 3, 4, 5}

	var _ [5]int

	_ = [...]int{1, 2, 3, 4, 5}

	oneTofive := [...]int{1, 2, 3, 4, 5}

	fmt.Println(oneTofive)

        oneTofive_cpy := oneTofive // copy of the array  
        oneTofive_cpy[0] = 100 
        println(oneTofive[0]==oneTofive_cpy[0]) 


        //slicing
        
}

func testCaseGo() (x int, y uint) {
	x = 10
	y = 20
	return

}
