package main

import "fmt"

func TSel() {
	fmt.Println("Welcome to the Selection function.")

	iffunc()
	ifelsefunc()
	nestedIfElse()

}

func iffunc() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	}

}

func ifelsefunc() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

}

func nestedIfElse() {

	/* local variable definition */
	var a int = 100
	var b int = 200

	/* check the boolean condition */
	if a == 100 {
		/* if condition is true then check the following */
		if b == 200 {
			/* if condition is true then print the following */
			fmt.Printf("Value of a is 100 and b is 200\n")
		}
	}
	fmt.Printf("Exact value of a is : %d\n", a)
	fmt.Printf("Exact value of b is : %d\n", b)

}
