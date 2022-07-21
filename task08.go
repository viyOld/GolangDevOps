//

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Result")

	fmt.Println(PartList([]string{"cdIw", "tzIy", "xDu", "rThG"}))
	fmt.Println(PartList([]string{"I", "wish", "I", "hadn't", "come"}))

}

func PartList(arr []string) string {
	l := len(arr)
	fin := ""
	for i := 0; i < l-1; i++ {
		fin += "("
		for j := 0; j < l; j++ {
			fin += arr[j]
			if i == j {
				fin += ","
			}
			if j < l-1 {
				fin += " "
			}
		}
		fin += ")"
	}
	return fin
}
