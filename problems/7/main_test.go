package main

import (
    "testing"
)

var primes = []int64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53}

func _testFindNPrimes(t *testing.T, fn func(int)(int64) ) {

	for i:=range primes {
		nstprime := fn(i+1)
		if nstprime != primes[i] {
			t.Errorf("%d-st prime should be %d, not %d", i+1, primes[i], nstprime)
		}
	}
}

func TestFindNPrimesNaive(t *testing.T) {
	_testFindNPrimes(t, nstPrimeNaive)
}

func TestFindNPrimesNaive2(t *testing.T) {
	_testFindNPrimes(t, nstPrimeNaive2)
}

func TestFindNPrimes(t *testing.T) {
	_testFindNPrimes(t, nstPrime)
}


func _benchTestNPrimes(b *testing.B, fn func(int)(int64)){
	for i:= 0; i<b.N; i++ {
		fn(1000)
	}
}
	
func BenchmarkNstPrimesNaive(b *testing.B){
	_benchTestNPrimes(b, nstPrimeNaive)
}

func BenchmarkNstPrimesNaive2(b *testing.B){
	_benchTestNPrimes(b, nstPrimeNaive2)
}
func BenchmarkNstPrimes(b *testing.B){
	_benchTestNPrimes(b, nstPrime)
}
