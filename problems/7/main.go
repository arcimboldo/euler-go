// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10001st prime number?
// 

package main

import (
    "fmt"
	"math"
	"flag"
)

var nst *int = flag.Int("n", 10001, "N-st prime")

// This is from problem 3
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

func nstPrimeNaive(n int) int64 {
	// find all n prime numbers in a stupid way
	var i int64 = 0
	var curprime int64 = 2
	for i=1; n>0; i++ {
		if isPrime(i) {
			n--
			curprime = i
		}
	}
	return curprime
}

func nstPrimeNaive2(n int) int64 {
	// find all n prime numbers in a slightly less stupid way
	var i int64 = 0
	var curprime int64 = 2

	// this allows us to jump 2 by 2
	if n == 1 { return 2}
	n--
	for i=1; n>0; i+=2 {
		if isPrime(i) {
			n--
			curprime = i
		}
	}
	return curprime
}

func nstPrime(n int) int64 {
	// special case. This allows us to jump 2 by 2
	if n == 1 { return 2}
	
	// Create an array of prime numbers
	var primes = make([]int64, n)


	primes[0] = 2
	var curprime int64 = 0
	var i int
	// loop 2 by 2, keep record of all prime numbers and devide by all of them

	for curprime = 3; n>1; curprime += 2 {
		var isprime bool = true
		for i= 0; primes[i] != 0; i++ {
			if primes[i] != 0 && curprime % primes[i] == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			primes[i] = curprime
			n--
		}
	}
	return primes[len(primes)-1]
}

func main(){
	flag.Parse()
	fmt.Printf("%d-st prime is %d\n", *nst, nstPrimeNaive2(*nst))
}
