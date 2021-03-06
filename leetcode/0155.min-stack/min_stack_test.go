package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.GetMin())

	stack.Push(3)
	assert.Equal(t, 3, stack.Top())
	assert.Equal(t, 2, stack.GetMin())

	stack.Push(1)
	assert.Equal(t, 1, stack.Top())
	assert.Equal(t, 1, stack.GetMin())

	stack.Pop()
	assert.Equal(t, 3, stack.Top())
	assert.Equal(t, 2, stack.GetMin())

	stack.Pop()
	stack.Pop()
	stack.Pop()
	assert.Equal(t, 0, stack.Top())
	assert.Equal(t, 0, stack.GetMin())
}
