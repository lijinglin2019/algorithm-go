package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func Postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var results []int

	for _, n := range root.Children {
		results = append(results, Postorder(n)...)
	}

	return append(results, root.Val)
}

func PostorderByStack(root *Node) []int {
	if root == nil {
		return nil
	}
	var results []int
	stack := []*Node{root}

	for len(stack) != 0 {
		last := len(stack) - 1
		n := stack[last]
		stack = stack[:last]

		results = append(results, n.Val)
		for i := range n.Children {
			stack = append(stack, n.Children[i])
		}
	}

	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	return results
}
