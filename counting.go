package main

// countingSort sorts the slice using the counting sort algorithm.
func countingSort(slice []Customer, max int) []Customer {
	// Initialize the count slice.
	count := make([]int, max+1)

	// Count the number of occurrences of each item.
	for _, item := range slice {
		count[item.numPurchases]++
	}

	// Adjust each count to reflect the counts before it.
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Place the items in sorted order.
	sorted := make([]Customer, len(slice))
	for _, item := range slice {
		sorted[count[item.numPurchases]-1] = item
		count[item.numPurchases]--
	}

	return sorted
}
