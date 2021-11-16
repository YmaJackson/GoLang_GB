//lesson4
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(100)
	fmt.Println("\n unsorted set  \n", slice)
	Sort(slice)
	fmt.Println("\n sorted set  \n", slice)
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func Sort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
