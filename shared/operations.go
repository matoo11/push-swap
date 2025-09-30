package shared

// All operation names
var ValidOps = map[string]struct{}{
	"sa": {}, "sb": {}, "ss": {},
	"pa": {}, "pb": {},
	"ra": {}, "rb": {}, "rr": {},
	"rra": {}, "rrb": {}, "rrr": {},
}

func Apply(stacks *Stacks, operation string) error {
	switch operation {
	case "sa":
		swap(&stacks.StackA)
	case "sb":
		swap(&stacks.StackB)
	case "ss":
		swap(&stacks.StackA)
		swap(&stacks.StackB)
	case "pa":
		push(&stacks.StackB, &stacks.StackA)
	case "pb":
		push(&stacks.StackA, &stacks.StackB)
	case "ra":
		rotate(&stacks.StackA)
	case "rb":
		rotate(&stacks.StackB)
	case "rr":
		rotate(&stacks.StackA)
		rotate(&stacks.StackB)
	case "rra":
		reverseRotate(&stacks.StackA)
	case "rrb":
		reverseRotate(&stacks.StackB)
	case "rrr":
		reverseRotate(&stacks.StackA)
		reverseRotate(&stacks.StackB)
	default:
		return ErrInvalidOp
	}
	return nil
}

func swap(stack *[]int) {
	values := *stack
	if len(values) < 2 {
		return
	}
	values[0], values[1] = values[1], values[0]
}

func push(from, to *[]int) {
	source := *from
	if len(source) == 0 {
		return
	}
	value := source[0]
	*from = source[1:]
	*to = append([]int{value}, *to...)
}

func rotate(stack *[]int) {
	values := *stack
	length := len(values)
	if length < 2 {
		return
	}
	firstValue := values[0]
	copy(values[0:], values[1:])
	values[length-1] = firstValue
}

func reverseRotate(stack *[]int) {
	values := *stack
	length := len(values)
	if length < 2 {
		return
	}
	lastValue := values[length-1]
	copy(values[1:], values[0:length-1])
	values[0] = lastValue
}
