package main

import (
	"testing"
)

var itoarn = []int{123, 98}
var itoara = [][]int{{1,2,3},{9,8}}

func TestIntToArray(t *testing.T){
	for i:= range itoarn {
		tmp := intToArray(itoarn[i])
		for j:= range tmp {
			if tmp[j] != itoara[i][j] {
				t.Errorf("%d as array should be %v not %v", itoarn[i], itoara[i], intToArray(itoarn[i]))
			}
		}
	}
}

var palindromes = []int {4, 44, 12321, 123321, 9009, 906609}

var notpal = []int { 12, 90 }

func TestIsPalindrome(t *testing.T) {
	for i:= range palindromes {
		if ! isPalindrome(palindromes[i]){
			t.Errorf("%d should be palindrome but it's not", palindromes[i])
		}
	}
}


func TestIsNotPalindrome(t *testing.T) {
	for i:= range notpal {
		if isPalindrome(notpal[i]){
			t.Errorf("%d should NOT be palindrome but it is", notpal[i])
		}
	}
}

var largepal = []int {9009, 906609}
var largefact = [][]int {
	{91, 99},
	{913, 993},
}

func TestLarger(t *testing.T){
	for idx:= range largepal {
		n := largepal[idx]
		i := largefact[idx][0]
		j := largefact[idx][1]
		nx, ix, jx := largerPalindrome(len(intToArray(largefact[idx][0])))
		if n != nx {
			t.Errorf("%d (%d*%d) is not the larger palindrome, should be %d (%d*%d)", nx, ix, jx, n, i, j)
		}
	}
}

func BenchmarkLP1(b *testing.B){
	for i:= 0; i<b.N; i++ {
		largerPalindrome(3)
	}
}

func BenchmarkLP2(b *testing.B){
	for i:= 0; i<b.N; i++ {
		largerPalindrome2(3)
	}
}
