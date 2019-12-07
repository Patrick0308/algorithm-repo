package tree

type Heap struct {
	data  []int
	n     int
	count int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data:  make([]int, capacity+1),
		n:     capacity,
		count: 0,
	}
}

func (h *Heap) Insert(a int) bool {
	if h.count >= h.n {
		return false
	}
	h.count = h.count + 1
	h.data[h.count] = a
	i := h.count
	for {
		if i/2 > 0 && h.data[i] > h.data[i/2] {
			h.swap(i, i/2)
			i = i / 2
		} else {
			break
		}
	}
	return true
}

func (h *Heap) swap(a int, b int) {
	tmp := h.data[a]
	h.data[a] = h.data[b]
	h.data[b] = tmp
}
