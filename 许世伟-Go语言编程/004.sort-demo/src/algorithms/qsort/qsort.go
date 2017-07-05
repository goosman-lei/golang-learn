package qsort

func QuickSort(in []int) []int {
	qsortRecursive(in, 0, len(in)-1)
	return in
}

func qsortRecursive(in []int, left, right int) {

	middle := (left + right) / 2
	sigVal := in[middle]

	posM := middle
	posL := left
	posR := right
	for posL < posR {
		for in[posL] <= sigVal && posL < posM {
			posL++
		}
		in[posM] = in[posL]
		posM = posL

		for in[posR] > sigVal && posR > posM {
			posR--
		}
		in[posM] = in[posR]
		posM = posR
	}
	in[posM] = sigVal

	if posM-left > 1 {
		qsortRecursive(in, left, posM-1)
	}
	if right-posM > 1 {
		qsortRecursive(in, posM+1, right)
	}
}
