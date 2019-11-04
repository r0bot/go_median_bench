package main

import (
	"fmt"
	"go_median_test/pkg"
	"math/rand"
)

func main() {
	sl := pkg.SList{
		List: []int{1, 2, 3},
	}

	sl.Add(rand.Int())
	fmt.Printf("%v", sl.List)
}
