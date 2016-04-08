package bubblesort

import "fmt"
import "testing"
import "math/rand"
import "time"

func TestBubbleSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	valuesLen := r.Int() % 30
	values := make([]int, valuesLen)
	for i := 0; i < len(values); i++ {
		values[i] = r.Int() % 10
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
