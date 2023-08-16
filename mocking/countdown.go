package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	times     = 3
	finalWord = "Go!"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (configurablesleeper *ConfigurableSleeper) Sleep() {
	configurablesleeper.sleep(configurablesleeper.duration)
}


type RealSleeper struct{}

func (s *RealSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := times; i > 0; i-- {
		fmt.Fprint(w, fmt.Sprintf("%d\n", i))
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord+"\n")
}

func main() {
	r := ConfigurableSleeper{
		duration: time.Second,
		sleep: time.Sleep,
	}
	Countdown(os.Stdout, &r)
}
