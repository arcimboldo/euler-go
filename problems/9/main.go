// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
// a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.Find the product abc.
// 

package main

import (
    "fmt"
	"flag"	
)

var flagsum = flag.Int64("n", 1000, "Sum of triplet")

func isTriplet(a, b, c int64) bool {
	return (a < b) && (b < c) && a*a + b*b == c*c
}

func findTripletWithSum(n int64) (a,b,c int64) {
	// naive
	for a=1; a < n; a++ {
		for b= 1; b < n-a; b++ {
			for c=1 ; c <= n-a-b; c++ {
				if isTriplet(a,b,c) && a + b + c == n {
					return a,b,c
				}
			}
		}
	}
	return 0,0,0
}


func main(){
	flag.Parse()
	a, b, c := findTripletWithSum(*flagsum)
	if c == 0 {
		fmt.Printf("No triplets with sum %d\n", *flagsum)
	} else {
		fmt.Printf("%d^2 + %d^2 = %d^2 = %d\n", a, b, c, *flagsum)
	}
}
