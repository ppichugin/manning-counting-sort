package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeRandomSlice(numItems, max int) []Customer {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]Customer, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = Customer{
			id:           fmt.Sprintf("C%d", i+1),
			numPurchases: random.Intn(max),
		}
	}
	return slice
}

// PrintSlice prints at most numItems items.
func printSlice(slice []Customer, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// CheckSorted verifies that the slice is sorted.
func checkSorted(slice []Customer) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
			println("The slice is NOT sorted!")
			return
		}
	}
	println("The slice is sorted")
}
