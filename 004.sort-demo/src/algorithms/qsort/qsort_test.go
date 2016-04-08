package qsort

import "fmt"
import "testing"
import "math/rand"
import "time"

func TestQsort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	valuesLen := r.Int() % 30
	values := make([]int, valuesLen)
	for i := 0; i < len(values); i++ {
		values[i] = r.Int() % 10
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
