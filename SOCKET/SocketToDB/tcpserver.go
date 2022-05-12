package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	//gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
}

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "minh17620"
	DB_NAME     = "minh"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connect Success " + DB_NAME + " " + DB_HOST + ":" + DB_PORT)
	return db
}
func InitDB() *gorm.DB {
	db = ConnectDB()
	return db
}

func main() {
	//Listen
	//Accept
	// Handle connect
	users := make(chan user)
	db = InitDB()
	db.AutoMigrate(&user{})
	dstream, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstream.Close()

	for {
		con, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(con, db)
	}

}

func handle(con net.Conn, db *gorm.DB) {
	var u user
	data, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	s := strings.Split(data, " ")

	u.Username = s[0]
	u.Email = s[1]
	db.Create(&u)
	handle(con, db)
}
