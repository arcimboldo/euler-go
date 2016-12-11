// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
//
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
// 

package main

import (
    "fmt"
)

func sumsquare(n int64) int64 {
	// compute the sum 1^2 + 2^2 + .. n^2
	// or, in LaTeX:
	//
	// \sum_{i=1}^n i^2
	//
	// There is a nice formula for this, cfr https://en.wikipedia.org/wiki/Square_pyramidal_number
	return n*(n+1)*(2*n+1)/6
}

func squaresum(n int64) int64 {
	// compute (1+2+..+n)^2
	//
	// also cfr https://en.wikipedia.org/wiki/Triangular_number

	x := n*(n+1)/2
	return x*x	
}

func main(){
	fmt.Printf("Result is %d\n", squaresum(100) - sumsquare(100))
}
