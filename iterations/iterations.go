package iterations

func Repeat(toRepeat string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += toRepeat
	}
	return repeated
}
