// Select dùng để chọn nhiều lệnh gửi/nhận từ các kênh
//Nếu nhiều lệnh được sẵn sàng, select sẽ chọn ngẫu nhiên
// select chỉ thực thi lệnh nhận khi trong kênh có data đã được truyền vào, ko thực thi đc với kênh rỗng
// Select không được bỏ trống
// Việc sử dụng timeout cũng rất quan trọng đối với các chương trình kết nối với tài nguyên bên ngoài hoặc cần thời gian thực thi
//Tiến hành timeout sẽ giúp Go Routine dễ dàng thực thi với Chanels và Select
package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}

//default_case thực thi khi những case khác chưa ready
// Trong trường hợp các kênh còn lại rỗng, select sẽ thực thi default
func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}
func default_case() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}

func main() {
	/*
		output1 := make(chan string)
		output2 := make(chan string)
		go server1(output1)
		go server2(output2)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		}*/

	default_case()

}

// Trong ví dụ trên, Time sleep của server 2 là 3s, trong khi server1 là 6s.
// Vì vậy server2() được thực thi trước và có output2 trước
