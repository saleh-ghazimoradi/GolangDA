package stack

type ReverseStack struct {
	items []rune
}

func (r *ReverseStack) Push(item rune) {
	r.items = append(r.items, item)
}

func (r *ReverseStack) Pop() (rune, bool) {
	if len(r.items) == 0 {
		return 0, false
	}
	item := r.items[len(r.items)-1]
	r.items = r.items[:len(r.items)-1]
	return item, true
}

func (r *ReverseStack) ReverseString(input string) string {
	stack := ReverseStack{}

	for _, ch := range input {
		stack.Push(ch)
	}

	var reversed []rune
	for {
		if char, ok := stack.Pop(); ok {
			reversed = append(reversed, char)
		} else {
			break
		}
	}
	return string(reversed)
}

func NewReverseStack() *ReverseStack {
	return &ReverseStack{items: make([]rune, 0)}
}
