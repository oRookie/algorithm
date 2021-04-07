package main

import "log"

func main() {
	//构建链表
	head := &RList{
		data: 1,
		next: nil,
	}
	head2 := &RList{
		data: 2,
		next: nil,
	}
	index, index2 := head, head2

	for i := 2; i < 9; i++ {
		node, node2 := &RList{
			data: i,
		}, &RList{
			data: 2 * i,
			next: nil,
		}
		index.next, index2.next = node, node2
		index, index2 = node, node2
	}

	log.Println(head)
	//log.Println(head2)
	//node3 := mergeSortedList(head, head2)
	//log.Println(node3)
	//log.Println(checkCircle(head))
	head4 := deleteLastKth(head, 10)
	middle := findMiddleNode(head)
	log.Println(head4)
	log.Println(middle)
}

//1	 2	3	4	5	6	7	8	9

//反转
func Reverse(node *RList) {
	var (
		head, previous *RList
	)
	for node != nil {
		next := node.next
		if next == nil {
			head = node
		}
		node.next = previous
		previous = node
		node = next
	}
	log.Println(head)
}

//检测环
func checkCircle(node *RList) bool {
	if node == nil {
		return false
	}
	fast := node
	slow := node
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}

//有序链表合并
func mergeSortedList(a *RList, b *RList) *RList {
	var head *RList
	if a.data < b.data {
		head = a
		a = a.next
	} else {
		head = b
		b = b.next
	}
	node := head
	for a != nil && b != nil {
		if a.data < b.data {
			node.next = a
			a = a.next
		} else {
			node.next = b
			b = b.next
		}
		node = node.next
	}
	if a != nil {
		node.next = a
	}
	if b != nil {
		node.next = b
	}
	return head
}

type RList struct {
	data int
	next *RList
}

/**
|
1	2	3	4	5	6	7	8	9
	2	3	4	5	6	7	8	9
	|
*/
//删除倒数第K个节点
func deleteLastKth(list *RList, k int) *RList {
	fast := list
	for i := 1; fast != nil && i < k; i++ {
		fast = fast.next
	}
	if fast == nil {
		return list.next
	}
	slow := list
	var prev *RList
	for fast != nil {
		fast = fast.next
		prev = slow
		slow = slow.next
	}
	if prev == nil {
		list = list.next
	} else {
		prev.next = prev.next.next
	}
	return list
}

//求链表的中间节点
func findMiddleNode(list *RList) *RList {
	fast := list
	slow := list
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
