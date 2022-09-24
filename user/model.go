package user

import "fmt"

type Gender string

const (
	Male    Gender = "MALE"
	Female  Gender = "FEMALE"
	unknown Gender = "UNKNOQN"
)

var (
	users      map[int]string
	Users2     map[int]string
	UserGender = map[int]Gender{}
)

// struct
// tipe data bentukan
// dimana kita bisa membentuk
// type data baru yang berisikan
// type data lainnya
type User struct {
	ID      uint64
	Name    string
	POB     string
	DOB     string
	address string
}

// assign function to struct
func (u User) CallName() {
	fmt.Println("Hello ", u.Name)
}

// aku ingin mengassign
// address pada user

func (u User) GetAddress() string {
	return u.address
}

func (u *User) SetAddress(in string) {
	// u == variable yang kita definisikan
	// dari segi memory address
	u.address = in
}

// fungsi kepemilikan value struct
// -> dia tidak akan mengubah nilai dari struct itu
// fungsi kepemilikan pointer/address struct
// -> dia akan mengubah nilai dari struct itu

// aku ingin membuat
// struct lain dengan nama Teacher dan Student
// dimana Teacher dan Student akan memiliki
// properti yang sama dengan User

type Teacher User
type Student struct {
	User
}
