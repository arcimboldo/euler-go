package main

import (
	"testing"
)

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

var divisors = [][]int64{
	{2*7, 7},
	{3*11, 11},
	{2*5*13, 13},
	{19, 19},
	{19*23, 23},
	{2*19*23, 23},
}

func TestBiggerDivisor(t *testing.T){
	for i:= range divisors {
		big := biggestDivisor(divisors[i][0])
		if big != divisors[i][1] {
			t.Errorf("%d is not the biggest divisor of %d, should be %d",
				big, divisors[i][0], divisors[i][1])
		}
	}

}


func TestBiggerDivisorSlow(t *testing.T){
	for i:= range divisors {
		big := slowBiggestDivisor(divisors[i][0])
		if big != divisors[i][1] {
			t.Errorf("%d is not the biggest divisor of %d, should be %d",
				big, divisors[i][0], divisors[i][1])
		}
	}

}

func BenchmarkBiggest(b *testing.B){
	for i := 0; i < b.N; i++ {
		biggestDivisor(2*19*23)
	}
}

func BenchmarkBiggestSlow(b *testing.B){
	for i := 0; i < b.N; i++ {
		slowBiggestDivisor(2*19*23)
	}
}
