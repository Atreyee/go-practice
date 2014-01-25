package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string, ch chan MyResult)
}

type MyResult struct{
	Body string
    Urls [] string
    Err error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan MyResult, urls map[string]bool) {
    // TODO: Fetch URLs in parallel.
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    if depth <= 0 {
        return
    }
	go fetcher.Fetch(url, ch)
    v := <-ch
    if(urls[url] == true){
    	return
    }
    urls[url] = true
    
    if v.Err != nil {
        fmt.Println(v.Err)
        return
    }
    fmt.Printf("found: %s %q\n", url, v.Body)
    for _, u := range v.Urls {
        Crawl(u, depth-1, fetcher, ch, urls)
    }
    return
}

func main() {
    ch := make(chan MyResult)
    urls:= make(map[string]bool)
    Crawl("http://golang.org/", 4, fetcher, ch, urls)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string, ch chan MyResult) {
    if res, ok := f[url]; ok {
        ch <- MyResult{res.body, res.urls, nil}
    }
    ch <- MyResult{"",nil, fmt.Errorf("not found: %s", url)}
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
