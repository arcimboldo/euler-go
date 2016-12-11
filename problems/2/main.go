// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import (
	"fmt"
)


func main(){
	n0, n1 := 1, 2
	sum := 0
	for n1 <= 4e6 {
		n0, n1 = n1, n0+n1
		if n0%2 == 0 {
			sum += n0
		}
	}
	fmt.Printf("Last: %d\n", n0)
	fmt.Printf("Sum: %d\n", sum)

}
