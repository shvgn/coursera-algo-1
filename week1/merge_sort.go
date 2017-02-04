package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// merge merges two slices and returns the resulting sorted slice
// and the number of inversions in it
func merge(s1, s2 []int) ([]int, int) {
	var i, j, x1, x2, inversions int
	var s1finished, s2finished bool

	l1 := len(s1)
	l2 := len(s2)
	length := l1 + l2
	merged := make([]int, length, length)

	for i < l1 || j < l2 {
		if i < l1 {
			x1 = s1[i]
		} else {
			s1finished = true
		}

		if j < l2 {
			x2 = s2[j]
		} else {
			s2finished = true
		}

		if s1finished || (!s2finished && x1 > x2) {
			merged[i+j] = x2
			j++
			if !s1finished {
				inversions += l1 - i
			}
		} else if s2finished || x1 <= x2 {
			merged[i+j] = x1
			i++
		}
	}
	return merged, inversions
}

// merge_sort returns the sorted slice and the number of inversions in it
func mergeSort(s []int) ([]int, int) {
	if len(s) <= 1 {
		return s, 0
	}
	if len(s) == 2 {
		x1, x2 := s[0], s[1]
		inversions := 0
		if x1 > x2 {
			s[0], s[1] = x2, x1
			inversions = 1
		}
		return s, inversions
	}
	divider := len(s) / 2
	s1, inv1 := mergeSort(s[:divider])
	s2, inv2 := mergeSort(s[divider:])
	merged, inv3 := merge(s1, s2)
	inversions := inv1 + inv2 + inv3
	return merged, inversions
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	fpath := os.Args[1]
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(err.Error())
	}

	numStr := string(data)
	nums := make([]int, 0)

	for _, s := range strings.Fields(numStr) {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(n))
	}
	_, inversions := mergeSort(nums)
	fmt.Printf("Found %d inversions in %s\n", inversions, fpath)
}
