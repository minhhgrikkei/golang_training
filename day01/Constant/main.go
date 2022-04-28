package main

/* Constant dùng để khai báo những giá trị cố định
 */
import "fmt"

func main() {
	const (
		name    = "John"
		age     = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)
}
