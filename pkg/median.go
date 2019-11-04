package pkg

import (
	"sort"
)

//Median with sort
type MList struct {
	List []int
}

//Naive approach where we sort every time (Not efficient the bigger it gets)
func (ml *MList) CalcMedian() int {
	//Sort the numbers if needed so we can find median
	if !sort.IntsAreSorted(ml.List) {
		sort.Ints(ml.List) // sort the numbers
	}

	mNumber := len(ml.List) / 2

	if len(ml.List)%2 == 0 {
		return ml.List[mNumber]
	}

	return (ml.List[mNumber-1] + ml.List[mNumber]) / 2
}

//The function adds to the sorted list and returns the new median
func (ml *MList) Add(n int) int {
	ml.List = append(ml.List, n)
	return ml.CalcMedian()
}

//Sorted insert
type SList struct {
	List []int
}

func (sl *SList) CalcMedian() int {
	mNumber := len(sl.List) / 2

	if len(sl.List)%2 == 0 {
		return sl.List[mNumber]
	}

	return (sl.List[mNumber-1] + sl.List[mNumber]) / 2
}

//The function adds to the sorted list and returns the new median
func (sl *SList) Add(n int) int {
	i := sort.SearchInts(sl.List, n)
	//Append to the slice to increase capacity
	sl.List = append(sl.List, 0)
	//Shift the items from the index where we should insert onwards
	copy(sl.List[i+1:], sl.List[i:])
	//Assign the value at the correct index
	sl.List[i] = n
	//Calculate the median of the new slice
	return sl.CalcMedian()
}
