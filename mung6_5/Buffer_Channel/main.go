// Khi ta gửi dữ liệu vào trong 1 kênh, lệnh gửi sẽ bị chặn ở kênh đó cho tới khi có một goroutine khác đọc dữ liệu kênh đó ra và cũng như vậy ngược lại đối với lệnh đọc dữ liệu từ một kênh
// Với một kênh đệm có thể được tạo bằng cách truyền một tham số dung lượng bổ sung cho make,hàm chỉ định kích thước của bộ đệm.

// Cú pháp: 		ch := make(chan type, capacity)

// Một kênh đệm chỉ bị chặn lệnh gửi(write) khi kênh đệm đầy và ngược lại với lệnh nhận(read) khi kênh rỗng

// capacity trong hàm make phải lớn hơn 0

package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch) // lệnh đóng kênh,người nhận có thể sử dụng một biến bổ sung trong khi nhận dữ liệu từ kênh để kiểm tra xem kênh đã bị đóng hay chưa.

}

func main() {
	ch := make(chan int, 2) // tạo kênh động với dung lượng là 2, do đó 2 lệnh write đầu sẽ không bị chặn,mà chỉ chặn ở lệnh tiếp theo
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		//  v, ok := <- ch , biến ok để kiểm tra xem kênh đóng hay chưa
		fmt.Println("read value", v, "from ch") // Mỗi lần đọc cách nhau 2s, và lệnh nhận(read) chỉ bị chặn khi kênh rỗng.
		time.Sleep(2 * time.Second)             //Ở trong ví dụ này vì time.sleep của write và read bằng nhau nên lệnh read không bị chặn

	}
}
