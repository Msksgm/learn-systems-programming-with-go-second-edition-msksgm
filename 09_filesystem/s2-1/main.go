package main

import (
	"fmt"
	"io"
	"os"
)

// 新規作成
func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Nwe file content\n")
}

// 読み込み
func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 06666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Append content\n")
}

func main() {
	open()
	read()
	append()
}
