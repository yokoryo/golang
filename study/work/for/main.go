//golangでは、Cのwhile文もforで表現する
package main

import (
	"fmt"
)
func main() {
	// Go
	n := 0
	for n < 10 {
		fmt.Printf("n = %d\n", n)
		n++
	}
}

	// C
	//  int n = 0;
	//  while (n < 10) {
	//      printf("n = %d\n", n);
	//      n++;
	//  }

