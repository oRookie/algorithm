package test

import (
	"log"
	"testing"
)

//定义二叉树的结构

type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

func TestTree(t *testing.T) {
	nodeE := &TreeNode{data: "e", left: nil, right: nil}
	nodeD := &TreeNode{data: "d", left: nil, right: nil}
	nodeB := &TreeNode{data: "b", left: nodeD, right: nodeE}
	nodeC := &TreeNode{data: "c", left: nil, right: nil}
	nodeA := &TreeNode{data: "a", left: nodeB, right: nodeC}
	//1.二叉樹的最大深度
	log.Println(maxDeath(nodeA))
}

/*
*1.求二叉树的最大深度
* 也就是求它左右子树最大深度+1
 */
func maxDeath(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left, right := maxDeath(node.left), maxDeath(node.right)
	return max(left, right) + 1
}

/**
* 2.求二叉树的最小深度
 */
func minDeath(node *TreeNode) int {
	if node == nil {
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
