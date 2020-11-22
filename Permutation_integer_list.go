package main

import "fmt"


func Permutation(data []int, i int, length int) {
	
	if length == i {

		PrintSlice(data)

		return
	}

	for j := i; j < length; j++ {

		swap(data, i, j)

		Permutation(data, i+1, length)

		swap(data, i, j)
	}

}

   func PrintSlice(data []int){

		fmt.Println(data)

	}

	func swap(data []int, x int, y int) {

		data[x], data[y] = data[y], data[x]
	}

	func main() {

		var data [5]int

		for i := 0; i < 5; i++ {

			data[i] = i
	}
	Permutation(data[:], 0, 5)
	
}