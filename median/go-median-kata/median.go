package go_median_kata

import "sort"

func Median(numbers ...[]float64) float64 {
	return medianOfSlice(collectMedians(numbers))
}

func collectMedians(numbers [][]float64) (medians []float64) {
	for _, nums := range numbers {
		medians = append(medians, medianOfSlice(nums))
	}
	return medians
}

func medianOfSlice(numbers []float64) float64 {
	sort.Float64s(numbers)

	length := len(numbers)
	switch {
	case length == 1:
		return numbers[0]
	case isEvenLength(length):
		return medianOfEvenSlice(numbers)
	case isOddLength(length):
		return medianOfOddSlice(numbers)
	default:
		return 0
	}
}

func medianOfOddSlice(numbers []float64) float64 {
	return numbers[(len(numbers)-1)/2]
}

func medianOfEvenSlice(numbers []float64) float64 {
	length := len(numbers)
	return (numbers[length/2] + numbers[(length-1)/2]) / 2
}

func isEvenLength(length int) bool {
	return length%2 == 0
}

func isOddLength(length int) bool {
	return length%2 == 1
}
