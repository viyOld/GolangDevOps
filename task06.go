//Given two numbers and an arithmetic operator (the name of it, as a string), return the result of the two numbers having that operator used on them.
//a and b will both be positive integers, and a will always be the first number in the operation, and b always the second.
//The four operators are "add", "subtract", "divide", "multiply".
//A few examples:(Input1, Input2, Input3 --> Output)
//5, 2, "add"      --> 7
//5, 2, "subtract" --> 3
//5, 2, "multiply" --> 10
//5, 2, "divide"   --> 2.5
//Try to do it without using if statements!

package main

import "fmt"

func main() {

	fmt.Println("Result")
	fmt.Println(Arithmetic(8, 2, "add"))
	fmt.Println(Arithmetic(8, 2, "subtract"))
	fmt.Println(Arithmetic(8, 2, "multiply"))
	fmt.Println(Arithmetic(8, 2, "divide"))
}

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		//fmt.Println("OS X.")
		return a + b
	case "subtract":
		//fmt.Println("Linux.")
		return a - b
	case "multiply":
		//fmt.Println("Linux.")
		return a * b
	case "divide":
		//fmt.Println("Linux.")
		return a / b
	default:
		//fmt.Printf("%s.\n", os)
		return 0
	}
	return 0
}
