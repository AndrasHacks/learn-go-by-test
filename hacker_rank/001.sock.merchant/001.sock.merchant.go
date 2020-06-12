package main

func SockMerchant(n int, socks []int) int {
	maxIndex := getMax(socks, n)
	helper := make([]int, maxIndex+1)
	for _, sock := range socks {
		helper[sock]++
	}
	var sum int
	for _, help := range helper {
		sum += help / 2
	}
	return sum
}

func getMax(slice []int, length int) int {
	max := 0
	for i := 0; i < length; i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}
