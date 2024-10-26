package main

import (
	"errors"
	"fmt"
)

var metersPerLiter float64

func main() {

	metersPerLiter = 10.0
	paintForRoomOne, err := paintAndAreaCalculatorWithErrHandling(4.2, 3.0)
	fmt.Printf("%.2f liters paints needed \n", paintForRoomOne)
	paintForRoomTwo, err := paintAndAreaCalculatorWithErrHandling(4.2, -3.0)
	fmt.Printf("%.2f liters paints needed \n", paintForRoomTwo)
	if err != nil {
		fmt.Println(err)

	}

	name := "Tom"
	//Get address of variable using & "address of" operator
	address := &name
	fmt.Println(*address)
	updateName(address)
	fmt.Println(name)
	fmt.Println(*address)

}

// function with multiple return values
func paintAndAreaCalculator(width, height float64) (float64, float64) {
	area := width * height
	return area, area / metersPerLiter
}

// function with error handling and multiple return values
func paintAndAreaCalculatorWithErrHandling(width, height float64) (paint float64, err error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, errors.New(fmt.Sprintf("a height of %.2f is invalid", height))
	}
	area := width * height
	return area / metersPerLiter, nil
}

// Go is a "pass-by-value" language; function parameters receive a copy
// of the arguments from the function call

//Pointers - oh good old pointer
//Get address of variable using & "address of" operator
// Values that represtn the address of a variable are known as Pointers

//Type of a point is written with a * symbol followed by the type of the variable
// the pointer points to
//*int - point to integer

//Get value at a pointer by *Variable name - value at Variable
// * operator can also be used to update the value at the pointer
// *myIntPointer = 9

func updateName(x *string) {
	*x = "Bob"
}

// It's okay to return a pointer to a local variable
// function returns the address of the local variable
func returnLocalPointer() *int {
	i := 5
	return &i
}
