package main

import (
	"fmt"
)

type chuoi []int

func list(test_list []int, rank chan chuoi) {
	var list chuoi
	for i := 0; i < len(test_list); i++ {
		list = append(list, 0)
	}
	rank <- list
}
func sort_rank(sort_list []int, list chuoi, rank_list chan chuoi) {
	for i := 0; i < len(sort_list); i++ {
		for j := 0; j < i; j++ {
			if sort_list[i] >= sort_list[j] {
				list[i]++
			} else {
				list[j]++
			}
		}
	}
	rank_list <- list
	/*
		for _, s := range list {
			fmt.Printf("%v ", s)
			fmt.Printf("\n")
		}*/
}
func main() {
	x := []int{14, 6, 5, 8, 7}
	rank_list := make(chan chuoi)
	go list(x, rank_list)
	rank := <-rank_list
	go sort_rank(x, rank, rank_list)
	ranklist := <-rank_list
	var ketqua chuoi
	for i := 0; i < len(x); i++ {
		ketqua = append(ketqua, 0)
	}
	for i := 0; i < len(x); i++ {
		ketqua[ranklist[i]] = x[i]
	}
	for i, s := range ketqua {
		fmt.Printf("%v ", i)
		fmt.Println(s)
	}
}
