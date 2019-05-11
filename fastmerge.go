package fastmerge

import (
	"sync"
)

const stepThreshold = 8

func mergeTwoSortedIntSlices(v []int, part1Start, part2Start int) {
	if part1Start == part2Start {
		// nothing to sort?
		return
	}
	if part1Start > part2Start {
		part1Start, part2Start = part2Start, part1Start
	}
	if part2Start > len(v) || part1Start > len(v) {
		// nothing to sort
		return
	}
	n1 := part2Start - part1Start
	n2 := n1
	if part2Start+n2 > len(v) {
		n2 = len(v) - part2Start
	}
	if n2 == 0 {
		// nothing to sort
		return
	}

	// minimal case
	if n1 == 1 {
		if v[part1Start] > v[part2Start] {
			v[part1Start], v[part2Start] = v[part2Start], v[part1Start]
		}
		return
	}

	// common case
	i1, i2 := 0, 0
	a := make([]int, 0, n1+n2)
	for j := 0; j < n1+n2; j++ {
		if i2 == n2 || i1 < n1 && (v[part1Start+i1] < v[part2Start+i2]) {
			a = append(a, v[part1Start+i1])
			i1++
		} else {
			a = append(a, v[part2Start+i2])
			i2++
		}
	}
	copy(v[part1Start:part2Start+n2], a)
}

func waitGroupWrapper(wg *sync.WaitGroup, f func([]int, int, int), v []int, part1Start, part2Start int) {
	f(v, part1Start, part2Start)
	wg.Done()
}

// MergeSortIntSlice sorts array of ints v in-place using additional memory buffers
func MergeSortIntSlice(v []int) {
	step := 1
	for {
		wg := sync.WaitGroup{}
		for j := 0; j < len(v); j += (step << 1) {
			if step <= stepThreshold {
				mergeTwoSortedIntSlices(v, j, j+step)
			} else {
				wg.Add(1)
				go waitGroupWrapper(&wg, mergeTwoSortedIntSlices, v, j, j+step)
			}
		}
		wg.Wait()
		step <<= 1
		if step >= len(v) {
			break
		}
	}
}
