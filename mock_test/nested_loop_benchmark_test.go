package main

import (
	"testing"
)

func NestedLoopTraversal(listSize int) int {
	list := make([]int, listSize)

	for i := 0; i < listSize; i++ {
		list[i] = i + 1
	}

	sum := 0
	for i := 0; i < listSize; i++ {
		for j := 0; j < listSize; j++ {
			sum += list[i] + list[j]
		}
	}

	return sum
}

func BenchmarkNestedLoopTraversal(b *testing.B) {
	listSize := 100
	for i := 0; i < b.N; i++ {
		NestedLoopTraversal(listSize)
	}
}
