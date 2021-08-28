package main

import (
	"testing"
)

var tree *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func init() {
	tree = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
}

/**
寻找二叉树第二小的节点
涉及知识点：对于二叉树的任意节点x，x的值不大于其所有节点的值，即二叉树的根节点就是所有节点的最小值
如果找到其中的一个节点，其值大于根节点的值。
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。

更正式地说，root.val = min(root.left.val, root.right.val) 总成立。

给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。

ans := -1
	rootVal := root.Val
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		//ans的原因是后续如果有
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
*/
func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	rootVal := root.Val
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		//退出条件 1。node为空等于遍历到底了，肯定需要return 2。找到第二小的值，并且下一个值大于第二小的值
		if node == nil || ans != -1 && node.Val > ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func TestFindSecondMinimumValue(t *testing.T) {
	println(findSecondMinimumValue(tree))
}

/**
二叉树寻路
在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。

如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；

而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。

给你树上某一个节点的标号 label，请你返回从根节点到该标号为 label 节点的路径，该路径是由途经的节点标号所组成的。

输入：label = 14
输出：[1,3,4,14]

方法一：数学
我们先来研究一个简单的情形：二叉树的每一行都是按从左到右的顺序进行标记。此时二叉树满足以下性质：

根节点位于第 1 行；

第1行有1个节点  2^{i-1}

第 i 行有 2^{i-1} 个节点，最左边的节点标号为 2^{i-1}2
i−1
 ，最右边的节点标号为 2^i-12
i
 −1；

对于标号为 \textit{val}val 的节点，其左子节点的标号为 2 \times \textit{val}2×val，右子节点的标号为 2 \times \textit{val} + 12×val+1，当 \textit{val}>1val>1 时，其父节点的标号为 \lfloor \frac{\textit{val}}{2} \rfloor⌊
2
val
​
 ⌋。

对于给定节点的标号 \textit{label}label，可以根据上述性质得到从该节点到根节点的路径，将路径反转后，即为从根节点到标号 \textit{label}label 的节点的路径。

回到这题，对于偶数行按从右到左的顺序进行标记的情况，可以转换成按从左到右的顺序进行标记的情况，然后按照上述思路得到路径，只要对偶数行的标号进行转换即可。为了表述简洁，下文将按从左到右的顺序进行标记时的节点的标号称为「从左到右标号」。

首先找到标号为 \textit{label}label 的节点所在的行和该节点的「从左到右标号」。为了找到节点所在行，需要找到 ii 满足 2^{i-1} \le \textit{label} < 2^i2
i−1
 ≤label<2
i
 ，则该节点在第 ii 行。该节点的「从左到右标号」需要根据 ii 的奇偶性计算：

当 ii 是奇数时，第 ii 行为按从左到右的顺序进行标记，因此该节点的「从左到右标号」即为 \textit{label}label；

当 ii 是偶数时，第 ii 行为按从右到左的顺序进行标记，将整行的标号左右翻转之后得到按从左到右的顺序进行标记的标号，对于同一个节点，其翻转前后的标号之和为 2^{i-1} + 2^i - 12
i−1
 +2
i
 −1，因此标号为 \textit{label}label 的节点的「从左到右标号」为 2^{i-1} + 2^i - 1 - \textit{label}2
i−1
 +2
i
 −1−label。

得到标号为 \textit{label}label 的节点的「从左到右标号」之后，即可得到从该节点到根节点的路径，以及路径上的每个节点的「从左到右标号」。对于路径上的每个节点，需要根据节点所在行的奇偶性，得到该节点的实际标号：

当 ii 是奇数时，第 ii 行的每个节点的「从左到右标号」即为该节点的实际标号；

当 ii 是偶数时，如果第 ii 行的一个节点的「从左到右标号」为 \textit{val}val，则该节点的实际标号为 2^{i-1} + 2^i - 1 - \textit{val}2
i−1
 +2
i
 −1−val。

最后，将路径反转，即可得到从根节点到标号 \textit{label}label 的节点的路径。

*/

func pathInZigZagTree(label int) []int {
	return nil
}
