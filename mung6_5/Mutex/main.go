//Mutex được sử dụng để cung cấp cơ chế khóa để đảm bảo rằng chỉ có một Goroutine đang chạy đoạn mã quan trọng tại bất kỳ thời điểm nào để ngăn chặn các điều kiện cuộc đua xảy ra.
//Có hai phương pháp được định nghĩa trên Mutex là Lock và Unlock
package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1 // lệnh nằm giữa Lock và Unlock đảm bảo mỗi Goroutine chạy đoạn mã trên độc lập. Các Goroutine sau sẽ nhận giá trị trả về từ goroutine trước đó
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m) // Lưu ý :cần chuyển địa chỉ waitgroup và mutex vào trong hàm
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
