package main

import "log"

func Sum(vals []int64) int64 {
	var total int64

	for _, val := range vals {
		if val%1e5 != 0 {
			total += val
		}
	}

	return total
}

func main() {
	log.Println(Sum([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
