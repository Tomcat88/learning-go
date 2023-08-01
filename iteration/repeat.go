package iteration

const repeatCount = 5

func Repeat(c string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += c
	}
	return
}
