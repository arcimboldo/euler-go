// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
// 

package main

import (
    "fmt"
	"math"
	"flag"
)

var flagsum = flag.Int64("n", 2*1000*1000, "")

// took from problem 3
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

// inspired by problem 3
func sumPrime(n int64) int64 {
	// find all prime numbers smaller than n
	var i int64 = 0
	var cursum int64 = 2

	// this allows us to jump 2 by 2
	if n < 2 { return 0}
	if n == 2 { return 2}

	for i=1; i <= n; i+=2 {
		if isPrime(i) {
			cursum  += i
		}
	}
	return cursum
}


func main(){
	flag.Parse()
	fmt.Printf("Sum of all primes below %d is %d\n", *flagsum, sumPrime(*flagsum))
}
