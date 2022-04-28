package main

import "fmt"

/* Có 4 kiểu dữ liệu cơ bản là string, float , int và bool
Có 3 cách khai báo biến:

Cách 1 : var + <tên biến> + kiểu dữ liệu - Với cách này nếu không gán giá trị thì biến sẽ có giá trị mặc định
Cách 2 : var + <tên biến> =  <giá trị gán>
Cách 3 : <tên biến> := <giá trị gán>

Bên cạnh đó có thể khai báo nhiều biến cung lúc: Khi đố tất cả các giá trị này có cùng 1 kiểu dữ liệu
	var <var1>, <var2>,... + <kieu du lieu> = <gia tri 1>, <gia tri 2>, ...


Với cách 2 và 3 thì biến bawsrt buộc phải có giá trị gán để xác định kiểu dữ liệu

*/
func main() {
	var student1 string = "John" //string
	var student2 = "Jane"        //string
	x := 2                       //int

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a string // giá trị mặc định là: ""
	var b int    // giá trị mặc định là 0
	var c bool   // giá trị mặc định là false

	fmt.Println("String's default is ", a)
	fmt.Println("Int's default is ", b)
	fmt.Println("Bool's default is ", c)
}
