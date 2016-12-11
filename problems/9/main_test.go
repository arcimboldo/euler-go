package main

import (
    "testing"
)

var triplets = [][]int64{
	{3,4,5},
	{5,12,13},
	{8,15,17},
}

var nottriplets = [][]int64{
	{3,4,6},
	{5,12,14},
	{8,15,19},
	// This is a reverse triplet
	{5,4,3},
	// this is an invalid triplet
	{0,0,0},
}

func TestIsTriplet(t *testing.T){
	for i:=range triplets {
		a,b,c := triplets[i][0], triplets[i][1], triplets[i][2]
		if ! isTriplet(a,b,c) {
			t.Errorf("%v should be a triplet but isTriplet fails", triplets[i])
		}
	}
}

func TestIsNotTriplet(t *testing.T){
	for i:=range nottriplets {
		a,b,c := nottriplets[i][0], nottriplets[i][1], nottriplets[i][2]
		if isTriplet(a,b,c) {
			t.Errorf("%v should NOT be a triplet but isTriplet fails", triplets[i])
		}
	}
}

func TestFindTripled(t *testing.T){

	for i:=range triplets{
		ta, tb, tc := triplets[i][0], triplets[i][1], triplets[i][2]
		a, b, c := findTripletWithSum(ta + tb + tc)
			
		if a != ta || b != tb || c != tc {
			t.Errorf("%v should be the tripled returned by findTripletWithSum(%d), not [%d,%d,%d]",
				triplets[i], ta+tb+tc, a,b,c)
		}
	}
	
}

func BenchmarkTriplet(b *testing.B){
	for i:=0; i<b.N; i++ {
		findTripletWithSum(380)
	}
}
