package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	fmt.Println("This module has intro to conditions, loops and methods")

	demoMethodWithTimePackage()

	demoStringMethods()
}

func demoStringMethods() {
	broken := "G# r#cks!"

	replacer := strings.NewReplacer("#", "o")
	// strings.NewReplacer returns a *Replacer

	fmt.Println("type of replace is :", reflect.TypeOf(replacer))
	fixed := replacer.Replace(broken)
	// Replacer has a Replace method
	fmt.Println(fixed)
}

func demoMethodWithTimePackage() {
	fmt.Println("time package has a Time type and a function called Now")

	fmt.Println("Now function return a value of time.Time type")
	var now time.Time = time.Now()

	fmt.Println("current time is :", now)

	fmt.Println("time.Time type has many methods, for example, we can use the Year method to get the current year")

	year := now.Year()
	month := now.Month()
	weekday := now.Weekday()
	yearDay := now.YearDay()

	fmt.Println("current year is :", year)
	// current year is : 2024
	fmt.Println("current month is :", month)
	// current month is : April
	fmt.Println("current weekday is :", weekday)
	// current weekday is : Wednesday
	fmt.Println("current yearDay is :", yearDay)

	fmt.Println("Methods are functions that are associated with values of a particular type.")
}
