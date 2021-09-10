package datastruct

import "testing"

/**
void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
*/

type Queue struct {
	Val []int
}

func (this *Queue) push(x int) {
	this.Val = append(this.Val, x)
}

func (this *Queue) pop() int {
	if this.empty() {
		return -1
	}
	result := this.Val[0]
	this.Val = this.Val[1:]
	return result
}

func (this *Queue) peek() int {
	if this.empty() {
		return -1
	}
	return this.Val[0]
}

func (this *Queue) empty() bool {
	return len(this.Val) == 0
}

func TestQueue(t *testing.T) {
	this := NewMyQueue()
	this.push(1)
	println(this.pop())
	println(this.pop())
}

func  NewMyQueue() Queue {
	return Queue{}
}