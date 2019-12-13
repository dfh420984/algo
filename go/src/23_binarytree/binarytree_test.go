package _23_binarytree

import "testing"

func TestTreeNode(t *testing.T) {
	var values []int = []int{3, 2, 4, 4, 5, 10, 7, 6, 9, 8}
	var root *TreeNode
	for _, v := range values {
		root = Insert(root, v)
	}
	res1 := PreOrderTraversal(root)
	res2 := InOrderTraversal(root)
	res3 := PostOrderTraversal(root)
	t.Log(res1)
	t.Log(res2)
	t.Log(res3)
}
