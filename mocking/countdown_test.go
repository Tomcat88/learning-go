package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls []string
}

func (spysleeper *SpySleeper) Write(p []byte) (n int, err error) {
	spysleeper.Calls = append(spysleeper.Calls, "write")
	return len(p), nil
}

func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func TestCountdown(t *testing.T) {
	t.Run("sleep and write should be interleaved", func(t *testing.T) {
		spySleeper := &SpySleeper{}

		Countdown(spySleeper, spySleeper)

		/* want := `3
		2
		1
		Go!
		` */
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(spySleeper.Calls, want) {
			t.Errorf("got %v, but want %v", spySleeper.Calls, want)
		}
	})
	t.Run("countdown should print the correct numbers", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(&buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}

type SpyTime struct {
	durationSleep time.Duration
}

func (spytime *SpyTime) Sleep(duration time.Duration) {
	spytime.durationSleep = duration
}

func TestConigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spytTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spytTime.Sleep,
	}
	sleeper.Sleep()

	if spytTime.durationSleep != sleepTime {
		t.Errorf("got %v, but want %v", spytTime.durationSleep, sleepTime)
	}
}
