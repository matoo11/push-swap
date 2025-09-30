package shared

type Stacks struct {
	StackA []int
	StackB []int
}

func (stacks Stacks) Clone() Stacks {
	stackACopy := make([]int, len(stacks.StackA))
	stackBCopy := make([]int, len(stacks.StackB))
	copy(stackACopy, stacks.StackA)
	copy(stackBCopy, stacks.StackB)
	return Stacks{StackA: stackACopy, StackB: stackBCopy}
}

func IsAscending(values []int) bool {
	for index := 1; index < len(values); index++ {
		if values[index-1] > values[index] {
			return false
		}
	}
	return true
}
