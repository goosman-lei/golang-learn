package bubblesort

func BubbleSort(in []int) []int {
	inLen := len(in)
	if inLen <= 1 {
		return in
	}

	for i := inLen - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
	return in
}
