//Розв'яжіть завдання:
//Даний цілий масив nums, виведіть на екран true, якщо будь-яке значення
//зустрічається в масиві як мінімум двічі, і false, якщо кожен елемент
//різний.

package main

import (
	"fmt"
	"os"
)

func main() {
	//nums := []int{5, 6, 7, 3, 4, 6, 7, 8, 2}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				fmt.Println("True")
				os.Exit(0)
			}
		}
	}
	fmt.Println("False")
}
