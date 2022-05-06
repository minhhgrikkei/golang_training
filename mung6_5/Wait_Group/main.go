//WaitGroup được sử dụng để đợi một tập hợp các Goroutines hoàn thành việc thực thi. Điều khiển bị chặn cho đến khi tất cả các Goroutines hoàn tất quá trình thực thi.
// Nếu trì hoãn nhiều lần trong 1 vòng for, defer sẽ trả theo stack LIFO
package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //Bộ đếm được giảm dần bởi lệnh gọi đến wg.Done
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) //Vòng lặp cũng tạo ra 3 process Goroutines
		go process(i, &wg)
	}
	wg.Wait() // main Goroutine đợi cho đến khi bộ đếm trở thành 0
	fmt.Println("All go routines finished executing")
}
