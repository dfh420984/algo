package _23_binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//插入节点
func Insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		node = new(TreeNode)
		node.Val = val
		return node
	}
	if val < node.Val {
		node.Left = Insert(node.Left, val)
	} else {
		node.Right = Insert(node.Right, val)
	}
	return node
}

//前序遍历(先打印根节点->左子树->右子树)
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		//取出节点
		e := stack[len(stack)-1]
		//更新stack
		stack = stack[:len(stack)-1]
		//放入根节点
		res = append(res, e.Val)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

//中序遍历
func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := InOrderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, InOrderTraversal(root.Right)...)
	return res
}

//后序遍历
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		lres := PostOrderTraversal(root.Left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.Right != nil {
		rres := PostOrderTraversal(root.Right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.Val)
	return res
}
