package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 2)
	results := make(chan int, 100)

	// 开启 3 个 goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// close(jobs) // 关闭之后还可以从通道取值

	//输出结果，可以使用 channel 阻塞或者 wait 来阻塞主协成直到子协成运行完毕
	for r := 1; r <= 5; r++ {
		a := <-results
		fmt.Printf("result:%d\n", a)
	}
}
