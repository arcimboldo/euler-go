package main

import (
	"testing"
)

var products = [][]int64{
	{2, 2},
	{3, 2*3},
	{10, 2520}, // problem data
	{20, 232792560}, // problem answer
}
func TestProduct(t *testing.T){
	for i:= range products {
		res := product(products[i][0])
		if res != products[i][1]{
			t.Errorf("Error with %d, got %d, should be %d", products[i][0], res, products[i][1])
		}
	}
}

var gcds = [][]int64{
	{ 2, 2, 2},
	{ 2, 3, 1},
	{10, 2, 2},
	{20, 30, 10},
}

func TestGcd(t *testing.T){
	for i:= range gcds {
		res := gcd(gcds[i][0], gcds[i][1])
		if res != gcds[i][2] {
			t.Errorf("gcd(%d,%d) should be %d, got %d", gcds[i][0], gcds[i][1], gcds[i][2], res)
		}
	}
}
