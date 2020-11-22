package main

import "fmt"

func main() {
	
	fmt.Println("Enter the size of slice")
	
	var n int	
	fmt.Scanln(&n)

    arr := make([]int, n)
	
	fmt.Println("Enter the elements of slice:")
	for i := 0; i < n; i++ {

		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	
	if n > 5 {
	
		fmt.Println(arr[:5])
	
	}

}