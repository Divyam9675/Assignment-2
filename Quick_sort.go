package main

import "fmt"

func main() {
	arr := []int{10, 15, 12, 26, 25, 22}
	sort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func sort(arr *[]int, low int, upper int) {
	if low < upper {
		pi := partition(arr, low, upper)
		sort(arr, low, pi-1)
		sort(arr, pi+1, upper)
	}
}

func partition(arr *[]int, low int, upper int) int {
	pivot := (*arr)[upper]
	i := low - 1
	for j := low; j < upper; j++ {
		if (*arr)[j] < pivot {
			i++
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = temp

		}
	}
	temp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[upper]
	(*arr)[upper] = temp
	return i + 1
}
