package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i<len(data); i++ {
		sum = sum + data[i]
	}
	fmt.Println(sum)
}