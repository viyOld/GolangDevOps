//The drawing below gives an idea of how to cut a given "true" rectangle into squares 
//("true" rectangle meaning that the two dimensions are different).
//Can you translate this drawing into an algorithm?
//You will be given two dimensions
//a positive integer length
//a positive integer width
//You will return a collection or a string (depending on the language; Shell bash, PowerShell, Pascal and Fortran return a string) with the size of each of the squares.
//Examples in general form: (depending on the language)
//  sqInRect(5, 3) should return [3, 2, 1, 1]
//  sqInRect(3, 5) should return [3, 2, 1, 1]
//  You can see examples for your language in **"SAMPLE TESTS".**

package main

import "fmt"

func main() {
	fmt.Println(SquaresInRect(5, 3))
	fmt.Println(SquaresInRect(3, 5))
	fmt.Println(SquaresInRect(5, 5))
	fmt.Println(SquaresInRect(20, 14))
}

func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth || lng == 0 || wdth == 0 {
		return nil
	}
	result := []int{}
	for (lng >= 1) && (wdth >= 1) {
		if lng > wdth {
			result = append(result, wdth)
			lng = lng - wdth
		} else if lng < wdth {
			result = append(result, lng)
			wdth = wdth - lng
		} else if lng == wdth {
			result = append(result, lng)
			break
		}
	}
	return result

}
