package sort

type HeapSort struct {
}

func sort(a []int) {
	n := len(a) - 1
	buildHeap(a, n)
	k := n
	for k > 0 {
		swap(a, 0, k)
		k--
		heapify(a, 0, k)
	}
}

func buildHeap(a []int, n int) {
	for i := (n+1)/2 - 1; i >= 0; i-- {
		heapify(a, i, n)
	}

}

func heapify(a []int, start int, end int) {
	for {
		max := start
		if start*2+1 <= end && a[start] < a[start*2+1] {
			max = start*2 + 1
		}
		if start*2+2 <= end && a[max] < a[start*2+2] {
			max = start*2 + 2
		}
		if max == start {
			break
		}
		swap(a, max, start)
		start = max
	}
}

func swap(a []int, n1 int, n2 int) {
	tmp := a[n1]
	a[n1] = a[n2]
	a[n2] = tmp
}
