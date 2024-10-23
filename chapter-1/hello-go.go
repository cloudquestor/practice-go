package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("--demo hello --")
	hello()
	fmt.Println("--demo helloName --")
	helloName("John")
	fmt.Println("--demo sqrt --")

	fmt.Println("demo runes below ----")

	demoRunes()
	fmt.Println("demo boolean ----")

	demoBool()

	fmt.Println("-- dmeo comparison --")
	demoComparison()

	fmt.Println("-- dmeo reflect --")
	demoReflect()

	fmt.Println("-- dmeo variable --")
	demoVariable()

	fmt.Println("-- dmeo zero value --")
	demoZeroValue()

	fmt.Println("-- dmeo for area --")
	demoArea()

}

func namingRules() {
	// if name begins with capital letter it is exported
	// expored things >> variable | functions | types

	// if name begins with small letter it is unexported

}

func demoArea() {
	length := 9
	width := 12.5

	// area := length * width - won't work as the type is not same

	area := float64(length) * width

	fmt.Println("Area is ", area)

}

func demoZeroValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// variable declaration

func demoVariable() {

	var name string
	var age int
	var salary float64

	// name = "Amit Singh"
	// age = 30
	// salary = 24750.49

	name, age, salary = "amit singh", 32, 2799.67

	fmt.Println("Name:", name)
	fmt.Println("Age", age)
	fmt.Println("Salary", salary)

	// declare and initialize or assign value at the same time
	var city string = "New York"

	//declare and assign value without type specification
	var state = "NY"

	//declare and assign value without type specification
	country := "USA"

	fmt.Println(city, state, country)

	//declare without var keyword and type - short variable declaration
	address := "123 Main St, Anytown USA 12345"
	fmt.Println(address)

}

// function example
func hello() {
	fmt.Println("Hello, World!")
}

// function with parameters
func helloName(name string) {
	fmt.Println("Hello,", name)
}

// function return value example
func sqrt() {
	fmt.Println("Hello, World!")
	fmt.Println("The square root of 16 is", math.Sqrt(16))
	fmt.Println(strings.Title("hello, world!"))
}

// runes - single char representation
// runes literal are single quotes

func demoRunes() {
	fmt.Println('A')
	fmt.Println('\t')
	fmt.Println('\n')
	fmt.Println('\\')
}

//boolean

func demoBool() {
	fmt.Println(true)
	fmt.Println(false)

}

//demo comparison

func demoComparison() {
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
}

// demo reflect
func demoReflect() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, World!"))
}
