package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miniorm/geeorm"
	"miniorm/geeorm/log"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	dsn := "root:abc123456@tcp(127.0.0.1:3306)/tv_test?charset=utf8mb4&parseTime=True"
	engine, err := geeorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Error(err)
		return
	}
	defer engine.Close()
	s := engine.NewSession()
	ok := s.Model(&User{}).HasTable()
	fmt.Println(ok)
}
