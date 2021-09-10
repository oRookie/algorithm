package datastruct

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
root = [1,null,2,3]
*/
func preorderTraversal(root *TreeNode) (val []int) {
	var pre func(*TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		val = append(val, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return

}

func inorderTraversal(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}
	vals = append(vals, inorderTraversal(root.Left)...)
	vals = append(vals, root.Val)
	vals = append(vals, inorderTraversal(root.Right)...)
	return
}

var levels [][]int

func init() {
	levels = [][]int{}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return levels
	}
	count(root, 0)
	return levels
}

func count(node *TreeNode, level int) {
	if len(levels) == level {
		levels = append(levels, []int{})
	}
	levels[level] = append(levels[level], node.Val)
	if node.Left != nil {
		count(node.Left, level+1)
	}
	if node.Right != nil {
		count(node.Right, level+1)
	}
}

func TestLevelOrder(t *testing.T) {
	fmt.Printf("%v", levelOrder(&TreeNode{
		Val: 1,
	}))
}

/**
 	3
9		20
null null 15 17

*/
func maxDepth(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return maxDepth(root.Left)
	} else if root.Right != nil && root.Left == nil {
		return maxDepth(root.Right)
	} else if root.Left != nil && root.Right != nil {
		return max(maxDepth(root.Left), maxDepth(root.Right))
	}
	return 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
		1
	2		2
3		4 4		3
*/
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func reverse(p, q *TreeNode) {
	if p == nil || q == nil {
		return
	}
	if p != nil && q != nil {
		tmp := q.Val
		q.Val = p.Val
		p.Val = tmp
	}
	reverse(p.Left, q.Right)
	reverse(p.Right, q.Left)
}

func TestIinvertTree(t *testing.T) {
	LogJson(invertTree(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}))

}

func LogJson(i interface{}) {
	bs, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	log.Println(string(bs))
}

func A() interface{} {
	var a interface{}
	a = "bb"
	return &a
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	right := root.Right
	root = searchBST(root.Left, val)
	if root == nil {
		root = searchBST(right, val)
	}
	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	//val插入到左
	if root.Val >= val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	} else { //val插入到右
		if root.Right == nil {
			root.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}
	return root
}
