//n^3 + (n-1)^3 + ... + 1^3 = m

package main

import (
        "fmt"
        "math"
)

func main() {

        fmt.Println("Result")
        fmt.Println(FindNb(4183059834009))
        fmt.Println(FindNb(24723578342962))
}

func FindNb(m int) int {
        n := 0
        sum := 0
        for {
                if sum < m {
                        n++
                        sum = sum + int(math.Pow(float64(n), 3))
                } else if sum == m {
                        return n
                } else if sum > m {
                        return -1
                }
        }
}
