/*
Một  chương trình golang tối thiểu gồm khai báo tên package
import <package>
Cuối cùng là hàm main

*/
package main // tên package thường là main. Với những package khác tên main, khi chạy sẽ không cho ra kết quả, dùng với mục đích tạo module

import (
	"fmt"
) // đây là nơi khai báo những package có trong thư viện của golang

// Dưới đây là hàm main
func main() {
	// This is a comment
	fmt.Println("Hello World!")
}

/*
bug         bắt đầu một báo cáo lỗi
build       biên dịch các gói và phụ thuộc
clean       biên dịch các gói và phụ thuộc
run         biên dịch và chạy chương trình Go
test        Chạy hàm test
*/
