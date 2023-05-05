package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	// 使用递归的方式解决
	var help func(node *TreeNode)
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		help(node.Left)
		ans = append(ans, node.Val)
		help(node.Right)
	}
	help(root)
	return ans
}

func main() {
	node := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(&node))
}
