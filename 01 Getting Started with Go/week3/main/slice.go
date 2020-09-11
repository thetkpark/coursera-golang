package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var arr = make([]int, 3)
	var count = 0
	for {
		var input string
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		var integer, error = strconv.Atoi(input)
		if error == nil {
			if count >= 3 {
				arr = append(arr, integer)
			} else {
				fmt.Println(count)
				arr[count] = integer
				count++
			}
			var forSort = append([]int(nil), arr...)
			sort.Ints(forSort)
			fmt.Println(forSort)
		}

	}
}
