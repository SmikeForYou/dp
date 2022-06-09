package dp

type BinaryTree[T any] struct {
	val   *T
	left  *BinaryTree[T]
	right *BinaryTree[T]
}

func (bt *BinaryTree[T]) weight(initialWeight int) int {
	if bt.val != nil {
		initialWeight += 1
	}
	if bt.left != nil {
		initialWeight = bt.left.weight(initialWeight)
	}
	if bt.right != nil {
		initialWeight = bt.right.weight(initialWeight)
	}

	return initialWeight
}

func (bt *BinaryTree[T]) toArray(initial []T) []T {
	if bt.val != nil {
		initial = append(initial, *bt.val)
	}
	if bt.left != nil {
		initial = bt.left.toArray(initial)
	}
	if bt.right != nil {
		initial = bt.right.toArray(initial)
	}

	return initial
}

func (bt *BinaryTree[T]) Weight() int {
	w := bt.weight(0)
	return w
}

func (bt *BinaryTree[T]) Insert(val T) {
	if bt.val == nil {
		bt.val = &val
		return
	}
	if bt.left == nil {
		bt.left = &BinaryTree[T]{val: &val}
		return
	}
	if bt.right == nil {
		bt.right = &BinaryTree[T]{val: &val}
		return
	}
	if bt.left.Weight() <= bt.right.Weight() {
		bt.left.Insert(val)
	} else {
		bt.right.Insert(val)
	}
}

func (bt *BinaryTree[T]) ToArray() []T {
	return bt.toArray(make([]T, 0))
}

func (bt *BinaryTree[T]) String() string {
	return ""
}

func NewBinaryTree[T any](data ...T) *BinaryTree[T] {
	root := new(BinaryTree[T])
	for _, val := range data {
		root.Insert(val)
	}
	return root
}

func WalkByBiaryTree[T any](bt *BinaryTree[T], callback func(node *BinaryTree[T])) {

}
