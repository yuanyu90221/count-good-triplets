package count_good_triplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	lengthOfArr := len(arr)
	for i := 0; i <= lengthOfArr-3; i++ {
		for j := i + 1; j <= lengthOfArr-2; j++ {
			for k := j + 1; k <= lengthOfArr-1; k++ {
				if checkRange(arr[i]-arr[j], arr[j]-arr[k], arr[i]-arr[k], a, b, c) {
					count += 1
				}
			}
		}
	}

	return count
}

func checkRange(d1 int, d2 int, d3 int, a int, b int, c int) bool {
	range1 := abs(d1)
	range2 := abs(d2)
	range3 := abs(d3)
	return range1 <= a && range2 <= b && range3 <= c
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
