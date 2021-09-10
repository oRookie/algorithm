package datastruct

import (
	"fmt"
	"sort"
	"testing"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}

func TestHasCycle(t *testing.T) {
	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	node.Next.Next.Next = node
	println(hasCycle(node))
}

/**
合并两个有序链表
1 2 3 4
2 2 3 4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{Next: head}
	for tmp := root; tmp.Next != nil; {
		if head.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return root
}

func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	//这个为什么要放到中间，是因为达到条件以后需要继续往下搜寻
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/**
1	2	3	4	4
			t

0	0	0	0	0
t
*/
func deleteDuplicates(head *ListNode) *ListNode {
	for tmp := head; tmp != nil && tmp.Next != nil; {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

func mergeTwoSlice(s1, s2 []int) []int {
	s := sort.IntSlice{}
	s = append(s1, s2...)
	s.Sort()
	fmt.Printf("%v", s)
	return s
}

func TestMergeTwoSlice(t *testing.T) {
	println(mergeTwoSlice([]int{1, 3, 5}, []int{2, 4, 6}))
}

