package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Elements of myArray: ", myArray)

	// 基于数组创建一个数组切片
	fmt.Println("Elements of myArray[:5]: ", myArray[:5])
	fmt.Println("Elements of myArray[5:]: ", myArray[5:])
	fmt.Println("Elements of myArray[:]: ", myArray[:])

	// 直接创建切片
	fmt.Println("Elements of make([]int, 5): ", make([]int, 5))
	fmt.Println("Elements of make([]int, 5, 10): ", make([]int, 5, 10))
	fmt.Println("Elements of []int{1, 2, 3, 4, 5}: ", []int{1, 2, 3, 4, 5})

	// 切片快速遍历
	fmt.Print("slice iterate with range []int{1, 2, 3, 4, 5}: ")
	for _, v := range []int{1, 2, 3, 4, 5} {
		fmt.Print(v, " ")
	}
	fmt.Println()

	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice): ", len(mySlice))
	fmt.Println("cap(mySlice): ", cap(mySlice))
	mySliceAppended := append(mySlice, 1, 2, 3)
	fmt.Println("after append len(mySlice): ", len(mySlice))
	fmt.Println("after append cap(mySlice): ", cap(mySlice))
	fmt.Println("after append len(mySliceAppended): ", len(mySliceAppended))
	fmt.Println("after append cap(mySliceAppended): ", cap(mySliceAppended))

	mySliceCopy := make([]int, 3, 6)
	copy(mySliceCopy, mySlice)
	fmt.Println("Elements of mySliceCopy: ", mySliceCopy)
	fmt.Println("len(mySliceCopy): ", len(mySliceCopy))
	fmt.Println("cap(mySliceCopy): ", cap(mySliceCopy))
}
