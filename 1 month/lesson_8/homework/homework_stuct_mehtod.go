package main

import (
	"fmt"
	"strings"
)

type User struct {
	name     string
	age      int
	birthday string
}

type AllUser struct {
	users []User
}

func (u AllUser) nameFilter(name string) {
	for _, user := range u.users {
		hasPrefix := strings.HasPrefix(strings.ToUpper(user.name), strings.ToUpper(name))
		if hasPrefix{
			fmt.Println(user)
		}

	}

}

func main() {
	var allUser = AllUser{
		users: []User{
			{
				name:     "Asatbek",
				age:      23,
				birthday: "2000",
			},
			{
				name:     "Asilbek",
				age:      20,
				birthday: "25-04-2003",
			},
			{
				name:     "Maxmud",
				age:      25,
				birthday: "12-07-1997",
			},
			{
				name:     "Azim",
				age:      22,
				birthday: "13-10-2001",
			},
			{
				name:     "asRor",
				age:      44,
				birthday: "03-07-1979",
			},
			{
				name:     "mAftuna",
				age:      25,
				birthday: "23-09-1997",
			},
			{
				name:     "MashXura",
				age:      16,
				birthday: "12-05-2007",
			},
		},
	}
	// for _,user:=range users{

	// }
	allUser.nameFilter("ma")
	allUser.nameFilter("az")
	allUser.nameFilter("a")
	allUser.nameFilter("as")
	// fmt.Printf()
}
