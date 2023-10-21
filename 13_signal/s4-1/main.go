package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 最初の 10 秒は Ctrl+C で止まる
	fmt.Println("Accept Ctrl+C for 10 seconds")
	time.Sleep(time.Second * 10)

	// 可変長引数で任意の数のシグナルを設定可能
	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	// 次の 10 秒は Ctrl+C を無視する
	fmt.Println("Ignore Ctrl+C for 10 seconds")
	time.Sleep(time.Second * 10)
}
