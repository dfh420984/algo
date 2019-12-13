package _24_binary_search_tree

import "testing"

func TestBinarySearchTree(t *testing.T) {
	bs := &BinarySearchTree{}
	bs.Insert(13)
	bs.Insert(10)
	bs.Insert(16)
	bs.Insert(9)
	bs.Insert(11)
	bs.Insert(14)
	tree := bs.GetTree()
	res := bs.InOrderTraversal(tree)
	t.Log(res)
	t.Log("---------------------------")
	res2 := bs.Find(16)
	t.Log(res2)
	res3 := bs.Find(99)
	t.Log(res3)
	t.Log("---------------------------")
	bs.Delete(10)
	tree = bs.GetTree()
	res4 := bs.InOrderTraversal(tree)
	t.Log(res4)
}
