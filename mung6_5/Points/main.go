// Con trỏ là một biến dùng để lưu trữ ĐỊA CHỈ của một biến khác
// Giá trị mặc định của con trỏ khi không gán là nil
package main

import (
	"fmt"
)

func main() {
	b := 255        // Biến b
	var a *int = &b // Con trỏ a lưu trữ địa chỉ b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Printf("value of a is %v\n", *a) // Giá trị mà a trỏ tới

	// Có thể tạo con trỏ thông qua hàm new
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}

// Lưu ý không nên để con trỏ tới một mảng arr làm một đối số của một hàm. Thay vì dùng *T[n] ta dùng slice T[]
