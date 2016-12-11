package main

import (
    "testing"
)

// took from problem 3
var primes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23}
var notprimes = []int64{1, 4, 6, 9, 12, 15, 121}

func TestIspirme(t *testing.T){
	for i := range primes {
		if ! isPrime(primes[i]) {
			t.Errorf("%d should be prime but is not", primes[i])
		}
	}
	for i := range notprimes {
		if isPrime(notprimes[i]){
			t.Errorf("%d should not be prime but it is", notprimes[i])
		}
	}
}

var primesum = [][]int64{
	{1, 0},
	{2, 2},
	{3, 2+3},
	{5, 2+3+5},
	{10, 2+3+5+7},
	{11, 2+3+5+7+11},
}

func TestSumPrime(t *testing.T){
	for i:=range primesum {
		if primesum[i][1] != sumPrime(primesum[i][0]) {
			t.Errorf("Sum of primes lower than %d should have sum %d instead we got %d",
				primesum[i][0], primesum[i][1], sumPrime(primesum[i][0]))
		}
	}
}
