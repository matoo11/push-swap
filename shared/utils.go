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
