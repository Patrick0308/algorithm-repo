package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{7, 6, 2, 3, 5, 6, 2, 4, 9, 5, 1, 13, 7}
	InsertionSort(a)
	fmt.Println(a)
}
