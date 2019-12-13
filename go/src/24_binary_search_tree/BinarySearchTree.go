package _24_binary_search_tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	tree *Node
}

func (b *BinarySearchTree) GetTree() *Node {
	return b.tree
}

//二叉平衡树查找
func (b *BinarySearchTree) Find(data int) *Node {
	p := b.tree
	for p != nil {
		if data < p.data { //小于当前节点，在左子数查找
			p = p.left
		} else if data > p.data { //大于当前节点，在右子数查找
			p = p.right
		} else {
			return p
		}
	}
	return nil
}

//二叉树插入操作
func (b *BinarySearchTree) Insert(data int) {
	if b.tree == nil {
		b.tree = &Node{data, nil, nil}
		return
	}

	p := b.tree
	for p != nil {
		if data > p.data { //在右子树中插入
			if p.right == nil { //右子树为空，直接插入
				p.right = &Node{data, nil, nil}
				return
			}
			//不为空找到右叶子节点
			p = p.right
		} else {
			if p.left == nil { //左子树为空，直接插入
				p.left = &Node{data, nil, nil}
				return
			}
			//不为空找到左叶子节点
			p = p.left
		}
	}
}

//删除节点
func (b *BinarySearchTree) Delete(data int) {
	p := b.tree  //初始化指向根节点
	var pp *Node //pp记录的是p的父节点
	//开始查找
	for p != nil && p.data != data {
		pp = p
		if data > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}
	if p == nil { //没有找到
		return
	}

	//要删除的有两个子节点
	if p.left != nil && p.right != nil { //查找右子树中最小节点
		minp := p.right
		minpp := p             //minpp表示minp的父节点
		for minp.left != nil { //右子树的最小节点一定在左边
			minpp = minp
			minp = minp.left
		}
		p.data = minp.data //将minp的数据替换到要删除的节点中

		//将指针指向叶子节点和其父节点,下面代码将变成删除叶子节点
		p = minp
		pp = minpp
	}

	//要删除的节点是叶子节点或者仅有一个节点
	var child *Node //p的子节点
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	//删除的是根节点
	if pp == nil {
		b.tree = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}
}

//中序遍历
func (b *BinarySearchTree) InOrderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	if root.left == nil && root.right == nil {
		return []int{root.data}
	}
	res := b.InOrderTraversal(root.left)
	res = append(res, root.data)
	res = append(res, b.InOrderTraversal(root.right)...)
	return res
}
