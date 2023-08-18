
package main

import (
	"fmt"
	"sync"
)

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func concurrentFactorial(n int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return factorial(n)
}

func main() {
	num := 5
	fmt.Printf("Calculating factorial of %d using goroutines\n", num)

	var wg sync.WaitGroup

	wg.Add()
	go func() {
		result := concurrentFactorial(num, &wg)
		fmt.Printf("Factorial of %d is %d\n", num, result)
	}()

	wg.Wait()
}
