package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeWeight(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	tree := &BinaryTree[int]{}
	tree.left = &BinaryTree[int]{}
	tree.right = &BinaryTree[int]{}
	tree.left.left = &BinaryTree[int]{}
	tree.left.right = &BinaryTree[int]{}
	tree.right.left = &BinaryTree[int]{}
	tree.right.right = &BinaryTree[int]{}
	tree.val = &data[0]
	tree.left.val = &data[1]
	tree.right.val = &data[2]
	tree.left.left.val = &data[3]
	tree.left.right.val = &data[4]
	tree.right.left.val = &data[5]
	tree.right.right.val = &data[6]
	assert.Equal(t, 7, tree.Weight())
	assert.Equal(t, 3, tree.left.Weight())
}

func TestNewBinaryTree(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewBinaryTree(data...)
	assert.Equal(t, 1, *tree.val)
	assert.Equal(t, 2, *tree.left.val)
	assert.Equal(t, 3, *tree.right.val)
	assert.Equal(t, 4, *tree.left.left.val)
	assert.Equal(t, 6, *tree.left.right.val)
	assert.Equal(t, 5, *tree.right.left.val)
	assert.Equal(t, 7, *tree.right.right.val)
}

func TestTreeToArray(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	arr := NewBinaryTree(data...).ToArray()
	assert.Equal(t, data, Sort(arr, func(a, b int) bool {
		return a < b
	}, false))

}
