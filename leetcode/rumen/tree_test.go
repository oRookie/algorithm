package rumen

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

/**
输入:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析:
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
1,1,1
1,1,0
1,0,1
*/

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

//广度优先
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	curColor := image[sr][sc]
	if curColor == newColor {
		return image
	}
	image[sr][sc] = newColor
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		for j := 0; j < 4; j++ {
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) && image[mx][my] == curColor {
				image[mx][my] = newColor
				queue = append(queue, []int{mx, my})
			}
		}

	}
	return image
}

/**
广度优先搜索。每次搜索到一个方格时，如果其与初始位置的方格颜色相同，就将该方格加入队列，并将该方格的颜色更新，以防止重复入队。
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	if currColor == newColor {
		return image
	}
	n, m := len(image), len(image[0])
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor
	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		for j := 0; j < 4; j++ {
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == currColor {
				queue = append(queue, []int{mx, my})
				image[mx][my] = newColor
			}
		}
	}
	return image
}

/**
我们从给定的起点开始，进行深度优先搜索。每次搜索到一个方格时，如果其与初始位置的方格颜色相同，就将该方格的颜色更新，以防止重复搜索；如果不相同，则进行回溯。
*/
func floodFillDfs(image [][]int, sr int, sc int, newColor int) [][]int {
	curColor := image[sr][sc]
	if curColor != newColor {
		dfs(image, sr, sc, curColor, newColor)
	}
	return image
}

func dfs(image [][]int, x, y, color, newColor int) {
	if image[x][y] == color {
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
				dfs(image, mx, my, color, newColor)
			}
		}
	}
}

func TestFloodFile(t *testing.T) {
	var images [][]int
	images = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Printf("%v", floodFillDfs(images, 1, 1, 2))
}

/**
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

示例 1:

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestMmergeTrees(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	LogJson(mergeTrees(root1, root2))
}

func LogJson(i interface{}) {
	bs, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	log.Println(string(bs))
}

type SkipTreeNode struct {
	Val   int
	Left  *SkipTreeNode
	Right *SkipTreeNode
	Next  *SkipTreeNode
}

func TestSkipTreeNode(t *testing.T) {
	node := &SkipTreeNode{
		Val: 1,
		Left: &SkipTreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
			Next: &SkipTreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
		},
		Right: &SkipTreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
			Next:  nil,
		},
	}
	LogJson(node)
}
