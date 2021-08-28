package main

type heap struct {
	val   []int
	cap   int //最大数据个数
	count int //已经存储的个数
}

func NewHeap(cap int) *heap {
	return &heap{
		val:   []int{},
		cap:   cap,
		count: 0,
	}
}

func (h *heap) Insert(val int) {
	if h.count >= h.cap {
		return
	}
	h.count = h.count + 1
	h.val[h.count] = val
	i := h.count
	for i/2 > 0 && h.val[i] < h.val[i/2] {
		h.swap(h.val[i], h.val[i/2])
	}
}


func (h *heap) swap(i, j int) {
	h.val[i], h.val[j] = h.val[j], h.val[i]
}

