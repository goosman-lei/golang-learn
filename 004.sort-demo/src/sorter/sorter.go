package main

import "flag"
import "fmt"
import "os"
import "bufio"
import "strconv"
import "strings"

import "algorithms/qsort"
import "algorithms/bubblesort"

var inFile *string = flag.String("i", "infile", "File contains values for sorting")
var outFile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if inFile != nil {
		fmt.Println("infile =", *inFile, "outFile =", *outFile, "algorithm =", *algorithm)
	}

	inFp, err := os.Open(*inFile)
	defer inFp.Close()
	if err != nil {
		panic(err)
	}

	inReader := bufio.NewReader(inFp)
	inData := make([]int, 0, 100)
	for row, err := inReader.ReadString('\n'); err == nil; row, err = inReader.ReadString('\n') {
		intVal, _ := strconv.Atoi(strings.TrimRight(row, "\n"))
		inData = append(inData, intVal)
	}

	var outData []int
	switch *algorithm {
	case "qsort":
		outData = qsort.QuickSort(inData)
	case "bubblesort":
		outData = bubblesort.BubbleSort(inData)
	}

	outFp, err := os.OpenFile(*outFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer outFp.Close()
	if err != nil {
		panic(err)
	}

	outWriter := bufio.NewWriter(outFp)
	for _, intVal := range outData {
		outWriter.WriteString(strconv.Itoa(intVal) + "\n")
	}
	outWriter.Flush()
}
