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
		return
	}

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
