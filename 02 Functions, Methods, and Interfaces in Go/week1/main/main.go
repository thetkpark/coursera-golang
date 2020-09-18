package main

import "fmt"

func main() {
	var arr = make([]int, 0)

	var n int

	fmt.Print("Enter length of array: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		arr = append(arr, number)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	fmt.Println(arr)
}

func swap(sli []int, index1 int, index2 int) {
	temp := sli[index1]
	sli[index1] = sli[index2]
	sli[index2] = temp
}
