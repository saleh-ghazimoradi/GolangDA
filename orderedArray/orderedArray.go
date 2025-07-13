package orderedArray

type OrderedArray struct {
	Elements []int
}

func (o *OrderedArray) Add(value int) {
	index := o.BinarySearch(value)
	o.Elements = append(o.Elements[:index], append([]int{value}, o.Elements[index:]...)...)
}

func (o *OrderedArray) Remove(value int) bool {
	index := o.BinarySearch(value)
	if index >= len(o.Elements) || o.Elements[index] != value {
		return false
	}
	o.Elements = append(o.Elements[:index], o.Elements[index+1:]...)
	return true
}

func (o *OrderedArray) Find(value int) (int, bool) {
	index := o.BinarySearch(value)
	if index < len(o.Elements) && o.Elements[index] == value {
		return index, true
	}
	return 0, false
}

func (o *OrderedArray) BinarySearch(value int) int {
	left, right := 0, len(o.Elements)
	for left < right {
		mid := (left + right) / 2
		if o.Elements[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func (o *OrderedArray) Get(index int) (int, bool) {
	if index < 0 || index >= len(o.Elements) {
		return 0, false
	}
	return o.Elements[index], true
}

func (o *OrderedArray) Element() []int {
	return append([]int{}, o.Elements...)
}

func (o *OrderedArray) Size() int {
	return len(o.Elements)
}

func (o *OrderedArray) Clear() {
	o.Elements = make([]int, 0)
}

func NewOrderedArray() *OrderedArray {
	return &OrderedArray{
		Elements: make([]int, 0),
	}
}
