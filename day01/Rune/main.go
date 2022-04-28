package main

import "fmt"

/* Đối với ASCII, mỗi character cần 1 byte để xác định, nhưng với UTF-8 thì có thể cần tới 3 hoặc 4 byte, đó là lí do vì sao cùng số ký tự nhưng lenght của biến kiểu rune lại dài hơn
Rune trong Golang được sinh ra để dùng biểu diễn những giá trị của UTF-8

*/
func main() {
	a_rune := []rune("ｿｹｯﾄを作成する")
	a := len("ｿｹｯﾄを作成する")
	b := len("abcdefghi")
	c := 'ｿ'
	for i, s := range a_rune {
		fmt.Print(i)
		fmt.Printf(" %c\n", s)
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
