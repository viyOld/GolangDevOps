//Count the number of Duplicates
//Write a function that will return the count of distinct case-insensitive alphabetic 
//characters and numeric digits that occur more than once in the input string. 
//The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits7.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	//s1 := "scahsccdfh"
	//s1 := "abcde" //-> 0 # no characters repeats more than once
	//s1 := "aabbcde" //-> 2 # 'a' and 'b'
	//s1 := "aabBcde" //-> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
	//s1 := "indivisibility" //-> 1 # 'i' occurs six times
	//s1 := "Indivisibilities" //-> 2 # 'i' occurs seven times and 's' occurs twice
	//s1 := "aA11" //-> 2 # 'a' and '1'
	s1 := "ABBA" //-> 2 # 'A' and 'B' each occur twice

	fmt.Println(duplicate_count(s1))

}

func duplicate_count(s1 string) int {
	rm := make(map[rune]int)
	rs := []rune(s1)
	l := 0
	for _, r := range rs {
		rm[unicode.ToLower(r)]++
		if rm[unicode.ToLower(r)] == 2 {
			l++
		}

	}
	return l //Instead of returning '1', you have to return integer 'i' as answer of solution.
}
