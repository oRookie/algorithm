package rumen

import (
	"fmt"
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
