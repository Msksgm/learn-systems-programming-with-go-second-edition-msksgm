package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock() をペアで呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func main() {
	// sync.Mutex 構造体の変数宣言
	// 次の宣言をしてもポインタ型になるだけで動作は同じ
	// mutex := new(sync.Mutex)
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
