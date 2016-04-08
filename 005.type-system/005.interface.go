package main

import "fmt"
import "reflect"

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWritter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

type IReaderWritter interface {
	IReader
	IWritter
}

type File struct {
}

func (f *File) Read(buf []byte) (n int, err error) {
	n = 0
	err = nil
	return
}
func (f *File) Write(buf []byte) (n int, err error) {
	n = 0
	err = nil
	return
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	pos = 0
	err = nil
	return
}
func (f *File) Close() (err error) {
	err = nil
	return
}
func (f File) Mock() {
}

func main() {
	obj := new(File)

	dumpMethods(obj)
	dumpMethods(*obj)
}

func dumpMethods(obj interface{}) {
	objType := reflect.TypeOf(obj)
	fmt.Println(objType)
	for i := 0; i < objType.NumMethod(); i++ {
		fmt.Println("\t" + objType.Method(i).Name)
	}
}

/*
Notice:
	1. 接口继承规则和类无区别
	2. 运行时caller-obj若为指针类型, 可访问全部方法. (golang会自动创建对应的非指针caller版本的函数). 但是, 如果caller-obj是非指针类型, 则只能访问定义中指定caller为非指针的方法. (比如上面代码*obj dumpMethods结果只有Mock)
*/
