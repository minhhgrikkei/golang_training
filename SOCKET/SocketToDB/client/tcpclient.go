package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)

	}

	var users []user
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		data := user.Username + " " + user.Email + "\n"
		fmt.Println(data)
		con.Write([]byte(data))
	}
	con.Close()
}
