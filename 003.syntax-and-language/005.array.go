package main

import "math/rand"
import "fmt"
import "time"

func main() {
	var arr_uint32 [8]uint32
	var arr_byte [8]byte
	var arr_struct [2][3]struct{ x, y int32 }
	var float64_a, float64_b, float64_c float64
	var arr_float64_ptr [3]*float64

	fmt.Println("byte && uint32 array")
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr_byte); i++ {
		arr_uint32[i] = rand.Uint32() % 256
		arr_byte[i] = byte(arr_uint32[i] % 256)
	}
	for i := 0; i < len(arr_byte); i++ {
		fmt.Printf("%02X %08X\n", arr_byte[i], arr_uint32[i])
	}
	fmt.Println()

	fmt.Println("struct 2d-array")
	for i := 0; i < len(arr_struct); i++ {
		for j := 0; j < len(arr_struct[i]); j++ {
			arr_struct[i][j].x = int32(i)
			arr_struct[i][j].y = int32(j)
		}
	}
	fmt.Println(arr_struct)
	fmt.Printf("%2s", "")
	for i := 0; i < len(arr_struct[0]); i++ {
		fmt.Printf(" | %14d", i)
	}
	fmt.Println("|")
	for i := 0; i < len(arr_struct); i++ {
		fmt.Printf("%2d", i)
		for j := 0; j < len(arr_struct[i]); j++ {
			fmt.Printf(" | x = %1d && y = %1d", arr_struct[i][j].x, arr_struct[i][j].y)
		}
		fmt.Println("|")
	}
	fmt.Println()

	fmt.Println("float64 array")
	float64_a = 0.1
	float64_b = 0.2
	float64_c = 0.3
	arr_float64_ptr[0] = &float64_a
	arr_float64_ptr[1] = &float64_b
	arr_float64_ptr[2] = &float64_c
	for i := 0; i < len(arr_float64_ptr); i++ {
		fmt.Printf("%2d %08X %0.2f\n", i, arr_float64_ptr[i], *arr_float64_ptr[i])
	}
	fmt.Println()

	fmt.Println("defined and initialized int array")
	arr_int_autoinited := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr_int_autoinited); i++ {
		fmt.Printf("%3d", arr_int_autoinited[i])
	}
	fmt.Println("\n")
}
