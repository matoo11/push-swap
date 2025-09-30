package shared

import (
	"strconv"
	"strings"
)


func ParseArgs(args []string) ([]int, error) {
	tokens := make([]string, 0, len(args))
	for _, argument := range args {
		for _, token := range strings.Fields(argument) {
			tokens = append(tokens, token)
		}
	}
	if len(tokens) == 0 {
		return []int{}, nil
	}

	parsedNumbers := make([]int, 0, len(tokens))
	seenValues := make(map[int]struct{}, len(tokens))
	for _, token := range tokens {
		value64, err := strconv.ParseInt(token, 10, 32)
		if err != nil {
			return nil, ErrInvalidNumber
		}
		value := int(value64)
		if _, exists := seenValues[value]; exists {
			return nil, ErrDuplicate
		}
		seenValues[value] = struct{}{}
		parsedNumbers = append(parsedNumbers, value)
	}
	return parsedNumbers, nil
}

func IndexOfMinimum(values []int) int {
	if len(values) == 0 {
		return -1
	}
	minimumIndex := 0
	minimumValue := values[0]
	for index := 1; index < len(values); index++ {
		if values[index] < minimumValue {
			minimumValue = values[index]
			minimumIndex = index
		}
	}
	return minimumIndex
}

func CompressToRanks(vals []int) []int {
	n := len(vals)
	if n == 0 {
		return []int{}
	}
	// Make a copy and sort
	sorted := make([]int, n)
	copy(sorted, vals)
	return compressUsingSort(vals)
}

// Split to keep imports simple
func compressUsingSort(vals []int) []int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	sortInts(tmp)

	rank := make(map[int]int, len(tmp))
	for index, value := range tmp {
		rank[value] = index
	}
	compact := make([]int, len(vals))
	for index, value := range vals {
		compact[index] = rank[value]
	}
	return compact
}

func sortInts(a []int) {
	if len(a) < 2 {
		return
	}
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
