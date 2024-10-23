package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This module has intro to conditions, loops and methods")

	demoMethodWithTimePackage()
}

func demoMethodWithTimePackage() {
	fmt.Println("time package has a Time type")

	var now time.Time = time.Now()

	fmt.Println("current time is :", now)
}
