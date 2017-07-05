package main

import "fmt"
import "os"
import "io"
import "os/user"

func main() {
	uinfo, _ := user.Current()
	fmt.Println("uinfo:", uinfo)

	if len(os.Args) != 3 {
		fmt.Printf("Usage:\n\t%s <file1> <file2>\n", os.Args[0])
		return
	}

	written, err := CopyFile(os.Args[2], os.Args[1])
	fmt.Println("written:", written, "err:", err)
}

func CopyFile(dst, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	defer srcFile.Close()
	if err != nil {
		panic(err)
		return
	}

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0755)
	defer dstFile.Close()
	if err != nil {
		panic(err)
		return
	}

	return io.Copy(dstFile, srcFile)
}

/*
Notice:
    1. defer语句应该是在函数体逻辑执行完成之后, 逆序向上执行的
    2. 仅会执行panic之前定义的defer
*/
