package test

import (
	"log"
	"math"
	"testing"
)

//定义二叉树的结构
type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

/**
二叉树分类：
1.满二叉树 所有节点都是满的
2.完全二叉树：除最后一层，其他所有节点都是满的，且最后一层只缺少右边若干的节点
3.平衡二叉树：任意节点的子树的最大高度差为1
*/

func TestTree(t *testing.T) {
	// nodeF := &TreeNode{data: "f", left: nil, right: nil}
	nodeE := &TreeNode{data: "e", left: nil, right: nil}
	nodeD := &TreeNode{data: "d", left: nil, right: nil}
	nodeB := &TreeNode{data: "b", left: nodeD, right: nodeE}
	// nodeC := &TreeNode{data: "c", left: nodeF, right: nil}
	nodeA := &TreeNode{data: "a", left: nodeB, right: nil}
	//1.二叉樹的最大深度
	log.Println("二叉树最大深度", maxDeath(nodeA))
	//2.二叉树的最小深度
	log.Println("二叉树最小深度", minDeath(nodeA))
	//3.二叉树的节点之和
	log.Println("二叉树的节点之和", numOfTreeNode(nodeA))
	//4.第k层的节点数
	log.Println("二叉树第k层节点个数", numOfLevelTreeNode(nodeA, 3))
	//5.判断平衡二叉树
	log.Println("判断平衡二叉树", isBalancedTree(nodeA))
}

/*
*1.求二叉树的最大深度
* 也就是求它左右子树最大深度+1
**/
func maxDeath(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left, right := maxDeath(node.left), maxDeath(node.right)
	return max(left, right) + 1
}

/**
* 2.求二叉树的最小深度
*
* 跟求最大深度思路一样，但是要注意，如果根节点的左子树或者右子树
**/
func minDeath(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right != nil {
		return minDeath(node.right) + 1
	}
	if node.right == nil && node.left != nil {
		return minDeath(node.left) + 1
	}
	return min(minDeath(node.left), minDeath(node.right)) + 1
}

/**
* 求二叉树中的节点个数
* 也就是求左右子树节点之和+1
 */

func numOfTreeNode(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return numOfTreeNode(node.left) + numOfTreeNode(node.right) + 1
}

/**
* 求二叉树中第k层的节点个数
* 第k层节点个数其实就是第k-1层根节点 左子树的节点数 和 右子树的节点数之和
*
*			1
*	2				3
*4		5		6
*
*k = 3
*numOfLevelTreeNode(node.left, 2) + numOfLevelTreeNode(node.right, 2)
*numOfLevelTreeNode(node.left.left, 1) + numOfLevelTreeNode(node.left.right,1) + numOfLevelTreeNode(node.right.left,1) + 0
*
 */
func numOfLevelTreeNode(node *TreeNode, k int) int {
	if node == nil || k < 1 {
		return 0
	}
	if k == 1 {
		return 1
	}
	return numOfLevelTreeNode(node.left, k-1) + numOfLevelTreeNode(node.right, k-1)
}

/**
* 判断二叉树是否是平衡二叉树
* 任何节点的子树高度差小于等于1（根节点的左子树和右子树的最大深度小于等于1）
 */
func isBalancedTree(node *TreeNode) bool {
	if node == nil {
		return true
	}
	left, right := maxDeath(node.left), maxDeath(node.right)
	if math.Abs(float64(left-right)) <= 1 {
		return isBalancedTree(node.left) && isBalancedTree(node.right)
	}
	return false
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
