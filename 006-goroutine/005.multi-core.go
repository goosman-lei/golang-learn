package main

import "os"
import "bufio"
import "fmt"
import "strconv"
import "strings"
import "runtime"

func main() {
	datas := readDatas()

	sum := 0
	sliceLen := len(datas) / 1
	for i := 0; i < 1; i++ {
		sum += doSum(datas[i*sliceLen : (i+1)*sliceLen])
	}

	fmt.Println(sum)
}

func readDatas() (datas []int) {
	dataFp, err := os.Open("./005.multi-core.data")
	defer dataFp.Close()
	if err != nil {
		panic(err)
	}

	dataReader := bufio.NewReader(dataFp)
	for row, err := dataReader.ReadString('\n'); err == nil; row, err = dataReader.ReadString('\n') {
		intVal, _ := strconv.Atoi(strings.TrimRight(row, "\n"))
		datas = append(datas, intVal)
	}
	return
}

func doSum(datas []int) (sum int) {
	for _, i := range datas {
		runtime.Gosched()
		sum += i
	}
	return
}
