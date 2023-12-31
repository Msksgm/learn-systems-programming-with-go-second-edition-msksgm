package main

import "time"

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutine が起動するときにはループが回りきって全部の task が最後のタスクになってしまう
			println(task)
		}()
	}
	time.Sleep(time.Second)
}
