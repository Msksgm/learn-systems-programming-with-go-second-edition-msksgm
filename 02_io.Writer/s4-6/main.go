package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

func main() {
	// 複数の io.Writer を受け取り、それらすべてに対して、書き込まれた内容を同時に書き込むデコレータ
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")

	// 書き込まれたデータを gzip 圧縮して、あらかじめ渡されていた os.File に中継するというサンプルコード
	file2, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer2 := gzip.NewWriter(file2)
	writer2.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer2.Close()

	// 出力結果を一時的にためておいて、ある程度の分量ごとにまとめて書き出す
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
