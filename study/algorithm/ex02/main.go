package main

import (
	"fmt"
)

func main() {
	data := []int{100, 90, 80, 70, 60}
	sum := 0
	for i := 0; i<len(data); i++ {
		sum = sum + data[i]
	}
	ave := sum / len(data)
	fmt.Println(ave)
}