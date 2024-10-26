package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	// fmt.Println("This module has intro to conditions, loops and methods")

	// demoMethodWithTimePackage()

	// demoStringMethods()

	// passFail()

	num := RandInt(15, 96)

	fmt.Println("Random number between 15 and 96 is:", num)

	for x := 0; x < 10; x++ {
		fmt.Println("x is :", x)
	}
}

func RandInt(min int, max int) int {
	// return min + rand.Intn(max-min)
	return rand.Intn(max)
}

func passFail() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	// input is a string
	fmt.Println("You entered:", input)
	// input is a string
	input = strings.TrimSpace(input)

	// err is only being assigned where as grade is declared and assigned
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	// grade is a string
	if grade >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}

	fmt.Println("The status is:", status)
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
