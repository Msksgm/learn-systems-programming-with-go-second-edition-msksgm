package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			// Lock() でロックしてからWait()を呼び出す
			mutex.Lock()
			defer mutex.Unlock()
			// Broadcast() が呼ばれるまで待つ
			cond.Wait()
			// 条件が満たされた
			println(name)
		}(name)
	}

	fmt.Println("よーい")
	time.Sleep(time.Second)
	fmt.Println("どん！")
	//待っているgoroutine を一斉に起こす
	cond.Broadcast()
	time.Sleep(time.Second)
}
