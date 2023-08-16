package main

import (
	"fmt"
	"testing"
)

// NestedLoopExample is a function that demonstrates nested loops.
func NestedLoopExample(rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Perform some operation here
			// For this example, we'll just print the loop indices
			fmt.Printf("Row: %d, Col: %d\n", i, j)
		}
	}
}

// BenchmarkNestedLoopExample measures the performance of the NestedLoopExample function.
func BenchmarkNestedLoopExample(b *testing.B) {
	// Run the function b.N times
	for i := 0; i < b.N; i++ {
		NestedLoopExample(10, 10) // Adjust rows and cols as needed
	}
}
