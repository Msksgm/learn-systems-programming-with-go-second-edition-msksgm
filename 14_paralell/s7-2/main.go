package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// ジョブ数をあらかじめ登録
	wg.Add(2)

	go func() {
		// 非同期で仕事をする（1）
		fmt.Println("仕事1")
		// Done() で完了を通知
		wg.Done()
	}()

	go func() {
		// 非同期で仕事をする（2）
		fmt.Println("仕事2")
		// Done() で完了を通知
		wg.Done()
	}()

	// 全ての処理が終わるまで待つ
	wg.Wait()
	fmt.Println("終了")
}
