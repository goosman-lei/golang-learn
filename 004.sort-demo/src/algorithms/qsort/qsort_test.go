package qsort

import "fmt"
import "testing"
import "math/rand"

func TestQsort(t *testing.T) {
	valuesLen := rand.Int() % 30
	values := make([]int, valuesLen)
	for i := 0; i < len(values); i++ {
		values[i] = rand.Int() % 10
	}

	fmt.Println("qsort nosort:", values)
	QuickSort(values)
	fmt.Println("qsort sorted:", values)
	for i := 0; i < len(values)-1; i++ {
		if values[i] > values[i+1] {
			t.Error("QuickSort() failed. Got", values)
		}
	}
}
