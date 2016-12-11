// A palindromic number reads the same both ways. The largest
// palindrome made from the product of two 2-digit numbers is 9009 =
// 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)


func intToArray(n int) []int {
	// find the biggest power of 10
	// split using n/i % 10 where i is from 10 to the biggest power 

	var max int = int(math.Log10(float64(n)))
	var array = make([]int, max+1)

	for i:=0; i<=max; i++ {
		array[i] = n / int(math.Pow(float64(10), float64(max-i))) % 10

	}
	
	return array
}

func isPalindrome(n int) bool{
	arrn := intToArray(n)
	j := len(arrn)-1
	for i:=0; i<=j; i++ {
		if arrn[i] != arrn[j] {
			return false
		}
		j--
	}
	return true
}


func largerPalindrome(digits int) (int, int, int) {
	imax := int(math.Pow(float64(10), float64(digits)))-1
	imin := int(math.Pow(float64(10), float64(digits-1)))
	biggest := 0
	curi, curj := 0, 0
	for i:= imax; i>=imin; i-- {
		for j:=imax; j>=i; j-- {
			// This might not be the biggest.
			if i*j > biggest {
				if isPalindrome(i*j){
					biggest = i*j
					curi = i
					curj = j
				}
			}
		}
	}
	return biggest, curi, curj
}

func largerPalindrome2(digits int) (int, int, int) {
	// This is incrediblly slower than the previous, can you guess why?
	imax := int(math.Pow(float64(10), float64(digits)))-1
	imin := int(math.Pow(float64(10), float64(digits-1)))
	biggest := 0
	curi, curj := 0, 0
	for i:= imax; i>=imin; i-- {
		for j:=imax; j>=i; j-- {
			// This might not be the biggest.
			if isPalindrome(i*j){
				if i*j > biggest {
					biggest = i*j
					curi = i
					curj = j
				}
			}
		}
	}
	return biggest, curi, curj
}

func main(){
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s n-digits\n", args[0])
		os.Exit(1)
	}
	digits, _ := strconv.Atoi(args[1])
	n, i, j := largerPalindrome(digits)
	fmt.Printf("n: %d i: %d, j: %d\n", n, i, j)

}
