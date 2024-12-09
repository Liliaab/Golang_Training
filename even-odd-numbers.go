package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, i := range arr {

		if i%2 == 0 {

			fmt.Println(i, "= even number")
		} else {
			fmt.Println(i, "=odd number")
		}

	}

}
