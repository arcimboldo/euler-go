// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
// 

package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int64) int64 {
	// find the greatest common divisor of a and b
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func product(n int64) int64 {
	// smallest number that is devided by all numbers from 1 to n
	var prod int64 = n
	for i:=n; i>1; i-- {
		if prod % i != 0 {
			prod = prod*i / gcd(prod, i)
		}
	}
	return prod
}

func main(){
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s n-digits\n", args[0])
		os.Exit(1)
	}
	max, _ := strconv.Atoi(args[1])
	fmt.Printf("Prod of first %d number is %d\n", max, product(int64(max)))
}
