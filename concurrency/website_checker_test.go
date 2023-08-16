package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func SlowWebsiteChecker(website string) bool {
	time.Sleep(20 * time.Millisecond)
	if website == "https://returnfalse.com" {
		return false
	} else {
		return true
	}
}

func MockWebsiteChecker(website string) bool {
	if website == "https://returnfalse.com" {
		return false
	} else {
		return true
	}
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"https://returnfalse.com",
		"https://www.google.com",
	}
	want := map[string]bool{
		"https://returnfalse.com": false,
		"https://www.google.com":  true,
	}
	got := CheckWebsites(MockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}

}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("https://www.website%d.com", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowWebsiteChecker, urls)
	}
}
