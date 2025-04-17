package Golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode, result ...[]int) []int {
	/*
	   Нужно обойти бинарное дерево в глубину.
	   Будем решать рекурсивной функцией
	*/
	if root == nil {
		return []int{}
	}
	if len(result) < 1 {
		result = append(result, []int{})
	}
	leftResult := inorderTraversal(root.Left, result[0])
	result[0] = append(result[0], leftResult...)

	result[0] = append(result[0], root.Val)

	rightResult := inorderTraversal(root.Right)
	result[0] = append(result[0], rightResult...)
	return result[0]
}
