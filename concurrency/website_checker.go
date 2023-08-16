package concurrency

import "time"

type WebsiteChecker func(website string) bool
type Result struct {
	string
	bool
}

func CheckWebsites(c WebsiteChecker, websites []string) map[string]bool {
	result := make(map[string]bool, len(websites))
	resultChan := make(chan Result)
	for _, website := range websites {
		go func(url string) {
			resultChan <- Result{url, c(url)}
		}(website)
	}

	for i := 0; i < len(websites); i++ {
		r := <-resultChan
		result[r.string] = r.bool
	}
	time.Sleep(2 * time.Second)
	return result
}
