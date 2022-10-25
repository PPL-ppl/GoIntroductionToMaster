package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Select(t *testing.T) {
	t.Run("Demo1", func(t *testing.T) {
		slowURL := "http://www.facebook.com"
		fastURL := "http://www.baidu.com"
		want := fastURL
		got, _ := select1(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("Http模拟服务器", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got := select2(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("chan", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := select3(serverA.URL, serverB.URL, 10)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
func select1(slowUrl, fastUrl string) (string, error) {
	startA := time.Now()
	http.Get(slowUrl)

	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(fastUrl)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return slowUrl, nil
	}

	return fastUrl, nil
}

func select2(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
func select3(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
