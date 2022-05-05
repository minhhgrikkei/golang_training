// Ticker sử dụng trong trường hợp ta muốn thực thi lại lệnh đó trong 1 khoảng thời gian
// Ở chương trình bên dưới, mỗi 500ms sẽ thực thi lệnh ở dòng 22 một lần
// Vì timeout là 1600ms cho go routine nên sẽ có 3 lần lệnh ở dòng 22 được thực thi

package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
