package main

import (
	"fmt"
	"strconv"

	"github.com/rexilyne/belajar-golang-3/user"
)

func main() {
	// conver string menjadi int
	num2, err := strconv.Atoi("1")
	fmt.Println(num2, err)
	// assignment struct
	// properti struct dengan huruf kecil
	// juga tidak akan bisa diakses oleh package lain
	var user1 user.User

	// cara 1
	user1.ID = 1
	user1.Name = "Test Name"
	user1.POB = "Indonesia"
	user1.DOB = "1960-08-08"
	fmt.Println(user1)

	user1 = user.User{
		ID:   10,
		Name: "Test Name 2",
		POB:  "Indonesia",
		DOB:  "1945-08-17",
	}
	fmt.Println(user1)
	user1.CallName()

	user2 := user.User{
		ID:   10,
		Name: "Test Name 3",
		POB:  "Indonesia",
		DOB:  "1945-08-17",
	}
	fmt.Println(user2)
	user2.CallName()

	// teacher dan student
	teacher1 := user.Teacher{
		ID:   99,
		Name: "Name99",
	}
	fmt.Println(teacher1)

	student1 := user.Student{}
	student1.Name = "Name100"
	fmt.Println(student1)
	student1.CallName()

	// apakah struct bisa memiliki
	// fungsinya dia sendiri?
	user3 := user2
	fmt.Println(user3.GetAddress())
	user3.SetAddress("Jakarta Selatan")
	fmt.Println(user3.GetAddress())

	// POINTER
	num := 9

	var num1 *int = &num // mengassign value dari pointer
	fmt.Println(num1, *num1)

	// pointer of struct
	ptrUser := &user.User{
		ID:   100,
		Name: "Tara",
	}
	fmt.Println(ptrUser, *ptrUser)
	ptrUser.SetAddress("hello")
	fmt.Println(ptrUser.GetAddress())

	usr := user.User{
		ID:   500,
		Name: "Lima Ratus",
	}
	usr.SetAddress("BSD")
	fmt.Println(usr.GetAddress())

	// function input parameter pointer
	cNum := 100
	ChangeNum(&cNum, 1000)
	fmt.Println(cNum)

	var ptrCNum *int = &cNum
	ChangeNum(ptrCNum, 10000)
	fmt.Println(*ptrCNum)

	// input param as pointer
	ChangeName(&usr, "Bzzz")
	fmt.Println(usr)

	// 1. init data student pada init function
	// 2. program akan mengeluarkan pesan error, ketika id tidak ditemukan
	// "student dengan id xx tidak ada pada database"
	// 3. id harus positive integer
}

func ChangeNum(source *int, destination int) {
	*source = destination
}

func ChangeName(src *user.User, name string) {
	(*src).Name = name
}
