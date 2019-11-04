package main

import (
	"go_median_test/pkg"
	"math/rand"
	"testing"
)

//Get a random int between 1 and 1000
func getRandInt() int {
	return rand.Intn(1000-1) + 1
}

var MedianResult int

func BenchmarkMedian(b *testing.B) {
	var r int
	ml := pkg.MList{
		List: []int{1, 2, 3},
	}
	for n := 0; n < b.N; n++ {
		r = ml.Add(getRandInt())
	}
	MedianResult = r
}

var SortedResult int

func BenchmarkSorted(b *testing.B) {
	var r int
	sl := pkg.SList{
		List: []int{1, 2, 3},
	}
	for n := 0; n < b.N; n++ {
		r = sl.Add(getRandInt())
	}
	SortedResult = r
}
