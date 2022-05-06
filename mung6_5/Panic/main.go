//Một số tình huống mà chương trình không thể tiếp tục thực hiện sau một tình trạng bất thường. Trong trường hợp này, sử dụng panic để kết thúc sớm chương trình.

package main

import (
	"fmt"
)

//Hàm này kiểm tra xem con trỏ firstName và lastName có nằm nil trong dòng không
////Vì chương trình kết thúc sau lệnh gọi hàm hoảng sợ ở dòng số. 12, mã ở dòng nos. 15, 17 và 18 sẽ không được thực thi.
/*
func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
*/
func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

//Nếu đó là nil các hàm gọi panic với một thông báo tương ứng. Thông báo này sẽ được in khi chương trình kết thúc.
func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

//Khi một hàm gặp sự cố, quá trình thực thi của nó sẽ bị dừng lại, mọi hàm hoãn lại được thực thi và sau đó điều khiển quay trở lại trình gọi của nó.
//Quá trình này tiếp tục cho đến khi tất cả các goroutine hiện tại return lại lúc chương trình in ra panic theo thứ tự stack rồi kết thúc
