package sum

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAll(arrays ...[]int) (sum []int) {
	for _, a := range arrays {
		sum = append(sum, Sum(a))
	}
	return
}

func SumAllTails(arrays ...[]int) (sum []int) {
	for _, a := range arrays {
		if a == nil || len(a) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(a[1:]))
		}
	}
	return
}
