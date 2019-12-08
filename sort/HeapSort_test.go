package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort_Sort(t *testing.T) {
	a := []int{7, 6, 2, 3, 5, 6, 2, 4, 9, 5, 1, 13, 7}
	sort(a)
	fmt.Println(a)
}
