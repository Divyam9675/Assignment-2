package main

import (
	"runtime"
	"sync"
	"fmt"
	
)


func main(){

		items := []int64{12,17, 9, 20, 63, 45, 19, 70, 100}

	  MergeSort(items)

	  fmt.Println(items)

}


func MergeSort(src []int64) {
	
	extraGoroutines := runtime.NumCPU() - 1
	semChan := make(chan struct{}, extraGoroutines)
	defer close(semChan)
	mergesort(src, semChan)
}


func mergesort(src []int64, semChan chan struct{}) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left := make([]int64, mid)
	right := make([]int64, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	wg := sync.WaitGroup{}

	select {
	case semChan <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left, semChan)
			<-semChan
			wg.Done()
		}()
	default:
	
		mergesort(left, semChan)
	}

	mergesort(right, semChan)

	wg.Wait()

	merge(src, left, right)
}

func merge(result, left, right []int64) {
	var l, r, i int

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}
}
