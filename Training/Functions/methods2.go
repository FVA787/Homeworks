package main

import (
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Salary int
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Salary: %d", u.Name, u.Age, u.Salary)
}

func main() {
	user := User{
		Name:   "Bob",
		Age:    20,
		Salary: 3050,
	}
	a := new(int)
	//Salary := 3000
	fmt.Println(user) // или  fmt.Println(user.String()), но в случае с методом String можно и просто fmt.Println(user).
	fmt.Println(*a)
	//fmt.Println(Salary)
	//fmt.Println(user.String())
}
