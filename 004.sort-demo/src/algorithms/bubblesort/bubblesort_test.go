package bubblesort

import "fmt"
import "testing"
import "math/rand"

func TestBubbleSort(t *testing.T) {
	valuesLen := rand.Int() % 30
	values := make([]int, valuesLen)
	for i := 0; i < len(values); i++ {
		values[i] = rand.Int() % 10
	}

	fmt.Println("bubblesort nosort:", values)
	BubbleSort(values)
	fmt.Println("bubblesort sorted:", values)
	for i := 0; i < len(values)-1; i++ {
		if values[i] > values[i+1] {
			t.Error("QuickSort() failed. Got", values)
		}
	}
}
