package utils

import (
	"fmt"
	"time"

	"github.com/mar-coding/goch01/app/models"
)

func merge(a []models.Str, ch chan models.Str) {
	// parallel merge sort
	if len(a) == 0 {
		close(ch)
		return
	}

	if len(a) == 1 {
		ch <- a[0]
		close(ch)
		return
	}

	mid := len(a) / 2

	ch1 := make(chan models.Str)
	go merge(a[:mid], ch1)

	ch2 := make(chan models.Str)
	go merge(a[mid:], ch2)

	v1, ok1 := <-ch1
	v2, ok2 := <-ch2

	for ok1 || ok2 {
		if (ok1 && ok2 && (v1.Greater(&v2) || v1.Equal(&v2))) || (ok1 && !ok2) {
			ch <- v1
			v1, ok1 = <-ch1
		} else if (ok1 && ok2 && v1.Lesser(&v2)) || (!ok1 && ok2) {
			ch <- v2
			v2, ok2 = <-ch2
		}
	}
	close(ch)
}

func MergeSort(a []models.Str) (sortedA []models.Str) {
	ch := make(chan models.Str)
	go merge(a, ch)

	for v := range ch {
		sortedA = append(sortedA, v)
	}
	return
}

func BubbleSort(inputs []models.Str) []models.Str {
	n := len(inputs)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if inputs[j].Lesser(&inputs[j+1]) {
				inputs[j].Swap(&inputs[j+1])
			}
		}
	}
	return inputs
}

func InsertSorted(slice []models.Str, item models.Str) []models.Str {
	// Find the index to insert 'item' on it in a way that the ordered slice stays ordered
	idx := searchIdx(slice, item)
	if idx == len(slice) {
		return append(slice, item)
	}
	slice = append(slice, models.Str{})
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = item
	return slice
}

func searchIdx(input []models.Str, item models.Str) int {
	// search for index to put data on that based on
	// binary search
	n := len(input)
	f := func(i int) bool {
		return input[i].Lesser(&item)
	}
	low, high := 0, n
	for low < high {
		mid := int((low + high) / 2) // avoid overflow when computing h
		// low ≤ mid < high
		if !f(mid) {
			low = mid + 1 // preserves f(i-1) == false
		} else {
			high = mid // preserves f(j) == true
		}
	}
	return low
}

func Benchmark(inputs []models.Str) {
	start1 := time.Now()
	_ = MergeSort(inputs)
	fmt.Printf("Time taken to sort with MergeSort Parallel: %v.\n", time.Since(start1))
	start2 := time.Now()
	_ = BubbleSort(inputs)
	fmt.Printf("Time taken to sort with BubbleSort: %v\n", time.Since(start2))
}
