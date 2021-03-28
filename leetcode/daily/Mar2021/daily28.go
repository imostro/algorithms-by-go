package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	root *trackTreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.root = build(root)
	return iterator
}

func (this *BSTIterator) Next() int {
	return 0
}

func (this *BSTIterator) HasNext() bool {
	return false
}

func build(root *TreeNode) (trackRoot *trackTreeNode)  {
	if root == nil {
		return nil
	}
	trackRoot = new(trackTreeNode)
	trackRoot.Val = root.Val
	trackRoot.Left = build(root.Left)
	trackRoot.Right = build(root.Right)

	return trackRoot
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type trackTreeNode struct {
	Val   int
	Left  *trackTreeNode
	Right *trackTreeNode
	isViewed bool
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
