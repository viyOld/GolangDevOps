//Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.
//Example
//CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
//The returned format must be correct in order to complete this challenge.
//Don't forget the space after the closing parentheses!

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Result")
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(numbers [10]uint) string {
	return "(" + strconv.Itoa(int(numbers[0])) + strconv.Itoa(int(numbers[1])) + strconv.Itoa(int(numbers[2])) + ") " +
		strconv.Itoa(int(numbers[3])) + strconv.Itoa(int(numbers[4])) + strconv.Itoa(int(numbers[5])) +
		"-" + strconv.Itoa(int(numbers[6])) + strconv.Itoa(int(numbers[7])) + strconv.Itoa(int(numbers[8])) + strconv.Itoa(int(numbers[9]))
}
