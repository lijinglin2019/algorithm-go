package leetcode

import (
	"testing"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		{},
		nil,
	}
	results := []int{3, 0, 0}

	for i := range cases {
		assert.Equal(t, results[i], DiameterOfBinaryTree(cases[i]))
	}
}