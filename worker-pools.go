package main

import (
	"fmt"
	"math"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- float64) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finised", j)
		results <- math.Pow(2, float64(j))
	}
}
func main() {
	const jobNum = 50
	jobs := make(chan int, jobNum)
	results := make(chan float64, jobNum)

	for w := 0; w < 50; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < jobNum; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 0; r < jobNum; r++ {
		res := <-results
		fmt.Println("job finished with", res)
	}

}
