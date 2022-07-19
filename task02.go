//Розв'яжіть завдання:
//Вам дано слайс складається з рядків. Необхідно надрукувати true тоді і тільки
//тоді, коли всі слова у слайсі відсортовані лексикографічно щодо
//один одного.

package main

import (
	"fmt"
	"os"
)

func main() {
	strs := []string{"abc", "xabd", "abf", "bdf", "fgh"}
	//strs := []string{"abc", "abd", "abf", "bdf", "fgh"}
	for i := 0; i < len(strs)-1; i++ {
		if strs[i] > strs[i+1] {
			fmt.Println("False")
			os.Exit(0)
		}
	}
	fmt.Println("True")
}
