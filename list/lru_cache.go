package main

import "log"

func main() {
	l := NewLRUCache(3)
	l.Get(1)
	l.Get(2)
	l.Get(1)
	l.Get(2)
	l.Get(3)
	l.Get(3)
	l.Get(1)
	l.Get(3)
	l.Get(5)
	l.Get(6)
	l.Get(7)
	node := l.Head
	for node.next != nil {
		log.Print(node.next.data,",","len: ",l.len)
		node = node.next
	}
}

type List struct {
	data interface{}
	next *List
}

type LRUCache struct {
	Cap  int
	len  int
	Head *List
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Cap: cap,
		Head: &List{
			data: nil,
			next: nil,
		},
	}
}

func (l *LRUCache) Get(val interface{}) bool {
	/**
	获取指定缓存数据
	思路：从链头开始遍历
	1.如果数据存在，遍历得到对应节点，并删除，然后重新插入头部
	2.如果数据不存在，则将新的数据插入到头部
		-如果链表容量已满，则删除链尾数据
	*/
	//记录尾节点前一个节点
	var prev *List
	node := l.Head
	for node.next != nil {
		if node.next.data == val {
			//数据存在 1。删除节点，重新插入头部
			l.del(node, node.next)
			newNode := &List{
				data: val,
				next: nil,
			}
			l.insertToHead(newNode)
			return true
		}
		prev = node
		node = node.next
	}
	//数据不存在
	newNode := &List{
		data: val,
		next: nil,
	}
	l.insertToHead(newNode)
	l.len++
	//如果长度超出容量，删除尾节点
	if l.len > l.Cap {
		l.del(prev, prev.next)
	}
	return false
}

func (l *LRUCache) insertToHead(node *List) {
	node.next = l.Head.next
	l.Head.next = node
}

func (l *LRUCache) del(prev *List, node *List) {
	prev.next = node.next
	node.next = nil
}
