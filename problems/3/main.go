 // The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

const tofactor int64 = 600851475143

func isPrime(n int64) bool {
	// Returns true if n is prime
	switch n {
	case 1:
		return false
	case 2:
		return true
	}
	var i int64 = 2
	for ; i<=int64(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func biggestDivisor(n int64) int64 {
	// Recursive function to find the biggest divisor of n.

	// It will find the smallest prime factor `p` of `n`, then call
	// itself on `n/p` and return the biggest of the two.

	// trivial case. This allows us to jump 2 by 2 in the for loop.
	if n == 2 { return 2}
	var i, biggest int64 = 3, 1
	for ; i<=n; i+=2 {
		if n % i == 0 && isPrime(i) {
			biggest = biggestDivisor(n/i)
			// I whish I could use math.Max, but
			// http://mrekucci.blogspot.ch/2015/07/dont-abuse-mathmax-mathmin.html
			if i > biggest {
				return i
			} else {
				return biggest
			}
		}
	}
	return n
	
}

func main(){
 	fmt.Printf("Biggest prime factor of %d: %d\n", tofactor, biggestDivisor(tofactor))

}

func slowBiggestDivisor(n int64) int64 {
	// return the biggest prime that devides n
	var i int64 = n
	for ; i > 1; i-- {
		if n % i == 0 && isPrime(i) {
			return i
		}
	}
	return n

}
