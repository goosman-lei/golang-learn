package main

import (
    "fmt"
    "bufio"
    "os"
    "compress/gzip"
    "time"
    "io"
)

func main() {
    gzipFile("example-12.7-gzipped.go")
    dgzipFile("example-12.7-gzipped.go.gz")
}

func gzipFile(fname string) {
    // 输出文件
    ofp, err := os.OpenFile(fname + ".gz", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Open output file failed: %T\n", err)
        os.Exit(1)
    }
    defer ofp.Close()

    // 输入文件
    ifp, err := os.Open(fname)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Open input file failed: %T\n", err)
        os.Exit(1)
    }
    defer ifp.Close()

    // gzip句柄
    zw := gzip.NewWriter(ofp)
    defer zw.Close()

    // setting header field
    zw.Name = fname + ".gz"
    zw.Comment = fname + " compressed with gzip"
    zw.ModTime = time.Now()

    fw := bufio.NewReader(ifp)
    for {
        bytes, isPrefix, err := fw.ReadLine()
        if err == io.EOF {
            break
        }
        zw.Write(bytes)
        if !isPrefix {
            zw.Write([]byte("\n"))
        }
    }

    zw.Flush()
}

func dgzipFile(fname string) {
    // 输入文件
    ifp, err := os.Open(fname)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Open input file failed: %T\n", err)
        os.Exit(1)
    }
    defer ifp.Close()

    zr, err := gzip.NewReader(ifp)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Create gzip Reader failed: %T\n", err)
        os.Exit(1)
    }
    defer zr.Close()

    fmt.Fprintf(os.Stderr, "ZipName: %s\n", zr.Name)
    fmt.Fprintf(os.Stderr, "ZipComment: %s\n", zr.Comment)
    fmt.Fprintf(os.Stderr, "ZipModTime: %s\n", zr.ModTime)

    buf := make([]byte, 10)
    for {
        rLen, err := zr.Read(buf)
        if err == io.EOF {
            return
        } else if err != nil {
            fmt.Fprintf(os.Stderr, "Read from gzip failed: %T\n", err)
            return
        }


        wLen := 0
        for wLen < rLen {
            tmpWLen, err := os.Stdout.Write(buf[wLen:rLen])
            if err != nil {
                fmt.Fprintf(os.Stderr, "Write to stdout err: %T\n", err)
                return
            }
            wLen += tmpWLen
        }
    }
}