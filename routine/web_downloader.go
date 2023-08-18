package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type DownloadResult struct {
	URL      string
	Duration time.Duration
	Err      error
}

func downloadPage(url string, wg *sync.WaitGroup, results []DownloadResult, index int) {
	defer wg.Done()

	startTime := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(startTime)

	result := DownloadResult{
		URL:      url,
		Duration: duration,
		Err:      err,
	}

	if err != nil {
		result.Err = fmt.Errorf("failed to download: %v", err)
	} else {
		resp.Body.Close()
	}

	results[index] = result
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.openai.com",
	}

	fmt.Println("Downloading web pages concurrently:")

	var wg sync.WaitGroup
	results := make([]DownloadResult, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go downloadPage(url, &wg, results, i)
	}

	wg.Wait()

	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("%s failed: %v\n", result.URL, result.Err)
		} else {
			fmt.Printf("%s took %s\n", result.URL, result.Duration)
		}
	}
}
