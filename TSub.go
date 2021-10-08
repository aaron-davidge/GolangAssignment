package main

import (
	"fmt"
)

func foo(test int) {
	test = 1 // change the value to 1
}

func bar(test *int) {
	*test = 1
}

func TSub() {
	fmt.Println("Welcome to the Subprogram function.")

	var i = 0
	fmt.Println("In main, i is", i)
	foo(i)
	fmt.Println("After foo call, in main i is", i)
	bar(&i)
	fmt.Println("After bar call, in main i is", i)

}
