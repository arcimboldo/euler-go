package main

import (
    "testing"
)

var square = [][]int64 {
	{2, 5},
	{3, 14},
	{4, 30},
	{10, 385},
}

func TestSumsquare(t *testing.T) {
	for i:= range square {
		if sumsquare(square[i][0]) != square[i][1] {
			t.Errorf("Sum of first %d^2 number should be %d, got %d",
				square[i][0], square[i][1], sumsquare(square[i][0]))
		}
	}
}

var ssquare = [][]int64 {
	{1, 1},
	{2, 9},
	{10, 3025},
	{100, 25502500},
}

func TestSquaresum(t *testing.T){
	for i:= range square {
		if squaresum(ssquare[i][0]) != ssquare[i][1] {
			t.Errorf("square of Sum of first %d number should be %d, got %d",
				ssquare[i][0], ssquare[i][1], squaresum(ssquare[i][0]))
		}
	}

}

