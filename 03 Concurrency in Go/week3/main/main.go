package main

import "fmt"

func sort(sli []int, start int, stop int, c chan []int) {
	for i := start; i <= stop; i++ {
		minIndex := i
		for j := i + 1; j <= stop; j++ {
			if sli[minIndex] > sli[j] {
				minIndex = j
			}
		}
		temp := sli[minIndex]
		sli[minIndex] = sli[i]
		sli[i] = temp
	}

	result := make([]int, 3)

	for i := 0; i < 3; i++ {
		result[i] = sli[start+i]
	}
	c <- result
}

func merge(sliL []int, sliR []int, c chan []int) {
	result := make([]int, 0)
	left, right := 0, 0

	for i := 0; i < len(sliL)+len(sliR); i++ {

		if left == len(sliL) {
			result = append(result, sliR[right])
			right++
		} else if right == len(sliR) {
			result = append(result, sliL[left])
			left++
		} else if sliL[left] <= sliR[right] {
			result = append(result, sliL[left])
			left++
		} else {
			result = append(result, sliR[right])
			right++
		}
	}
	c <- result
}

func main() {
	c := make(chan []int)
	arr := make([]int, 12)

	for i := 0; i < 12; i++ {
		fmt.Scan(&arr[i])
	}

	go sort(arr, 0, 2, c)
	go sort(arr, 3, 5, c)
	go sort(arr, 6, 8, c)
	go sort(arr, 9, 11, c)

	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c

	go merge(arr1, arr2, c)
	go merge(arr3, arr4, c)

	result1 := <-c
	result2 := <-c

	go merge(result1, result2, c)

	fmt.Println(<-c)
}
