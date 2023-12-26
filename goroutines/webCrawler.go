package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(URL string, depth int, fetcher Fetcher,
	visited map[string]bool,
	waitGroup *sync.WaitGroup,
	mutex *sync.Mutex) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	defer waitGroup.Done()
	mutex.Lock()

	// The current URL already has been investigated. Exit!
	if visited[URL] {
		mutex.Unlock()
		return
	}

	// The current URL hasn't been investigated yet. Now it will be crawled,
	// so set the visited boolean map as true and unlock the mutex
	visited[URL] = true
	mutex.Unlock()

	// If the recursion level reaches 0,
	// no need to investigate further
	if depth <= 0 {
		return
	}

	// Now we specified the URL will need to be fetched (once at least)
	body, URLs, err := fetcher.Fetch(URL)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", URL, body)

	var subWaitGroup sync.WaitGroup
	for _, URL := range URLs {
		subWaitGroup.Add(1)
		go func(URL string) {
			defer subWaitGroup.Done()
			waitGroup.Add(1)
			Crawl(URL, depth-1, fetcher, visited, waitGroup, mutex)
		}(URL)
	}
	subWaitGroup.Wait()

	return
}

func main() {
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	visited := make(map[string]bool) // Will log if the currently received URL has been investigated

	waitGroup.Add(1)
	Crawl("https://golang.org/", 4, fetcher, visited, &waitGroup, &mutex)
	waitGroup.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
