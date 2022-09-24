package user

import "fmt"

func GetUser(i int) string {
	return users[i]
}

func StoreUser(i int, name string) {
	users = map[int]string{}
	users[i] = name
}

func init() {
	for i := 0; i < 10; i++ {
		fmt.Println("print from init function in user", i)
	}
}
