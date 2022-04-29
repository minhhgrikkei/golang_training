package main

import "fmt"

func main() {
	/*
		a_string := "ｿｹｯﾄを作成する"
		a := len("ｿｹｯﾄを作成する")
		b := len("abcdefghi")
		c := 'ｿ'
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(string(a_string[0:3]))
		fmt.Println(c)
	*/
	input_string := "ｿｹｯﾄを作成するabcdefghi"

	ketqua := startandend(0, 10, input_string)
	for i, s := range ketqua {
		fmt.Print(i)
		fmt.Printf(" %c\n", s)
	}

}

type chuoi []rune

func startandend(start int, end int, input_string string) []rune {

	var ketqua chuoi
	var temp chuoi = []rune(input_string)
	ketqua = temp[start : end+1]

	return ketqua
}
