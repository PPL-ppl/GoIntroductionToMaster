package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_GoRouting(t *testing.T) {

	t.Run("", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		actualResults := CheckWebsites2(mockWebsiteChecker, websites)
		want := len(websites)
		got := len(actualResults)
		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
		expectedResults := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		if !reflect.DeepEqual(expectedResults, actualResults) {
			t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
		}
	})

}
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.N = 52
	for i := 0; i < b.N; i++ {
		CheckWebsites2(slowStubWebsiteChecker, urls)
	}
}

type WebsiteChecker func(string) bool

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	mp := make(map[string]bool)
	for _, url := range urls {
		mp[url] = wc(url)
	}
	return mp
}

type result struct {
	string
	bool
}

func CheckWebsites2(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
