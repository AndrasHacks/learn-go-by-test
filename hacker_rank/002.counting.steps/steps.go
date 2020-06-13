package main

func CountSteps(n int, steps string) int {
	level := 0
	countOfValleys := 0
	for i := 0; i < n; i++ {
		const down = 'D'
		if steps[i] == down {
			level--
		} else {
			level++
		}
		if level == -1 && steps[i] == down {
			countOfValleys++
		}
	}
	return countOfValleys
}
