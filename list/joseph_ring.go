package main

import "log"

/**
约瑟夫环
题目：有n个人围成一圈，顺序排号。从第一个人开始报数（从1到m报数），凡报到m的人退出圈子，直到船上仅剩 r人为止,问都有哪些编号的人下船了呢？。
假设n=30，m=9，r=15
*/
func main() {
	n, m, r := 30, 9, 15
	//构建环
	head := &JList{
		data: 1,
	}
	index := head
	for i := 2; i < n+1; i++ {
		node := &JList{
			data: i,
			next: nil,
		}
		index.next = node
		index = node
	}
	index.next = head

	//构建完成，开始报数
	var result []interface{}
	p := head
	count := 1
	for n != r {
		p = p.next
		count += 1
		if count == m {
			result = append(result, p.data)
			p.next = p.next.next
			count = 0
			n--
		}
	}
	log.Printf("出列顺序为：%v", result)
}

type JList struct {
	data interface{}
	next *JList
}
