package main

import "testing"

type testpair struct {
	value int
	number int
	sum int
}

var tests=[]testpair {
	{1234, 4, 10},
	{0,1,0},
	{-3,0,0},
	{55,2,10},
}

func TestCountLenght(t *testing.T) {
	for _,v:=range tests {
		temp:=CountLenght(v.value)
		if temp!=v.number {
			t.Error("For", v.value,
				"Expected", v.number,
				"Got", temp)
		}
	}
}

func TestSumNum (t *testing.T) {
	for _,v:=range tests {
		temp:=SumNum(v.value)
		if temp!=v.sum {
			t.Error("For", v.value,
				"Expected", v.sum,
				"Got", temp)
		}
	}
}