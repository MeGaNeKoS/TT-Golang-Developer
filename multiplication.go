package main

import "fmt"


func main() {
	max := 99
	for i := 1; i <= max; i++ {
		for j := 1; j <= max; j++ {
			result := fmt.Sprintf("%d * %d = %d ", i, j, i*j)
			fmt.Print(result)
			fmt.Print("\t")

		}
		fmt.Println()
	}
}
