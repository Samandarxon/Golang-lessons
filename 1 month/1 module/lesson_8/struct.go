package main

import "fmt"

type Human struct {
	// field
	firstName string
	lastName  string
	age       int
	width     float64
	height    float64
	birthday  string // 2000-05-29
}

func makeSingleUser() Human {

	// var user int
	// var user Human = Human{
	// 	"Asadbek",
	// 	23,
	// 	"Ergashev",
	// 	82,
	// 	171,
	// 	"2000-05-29",
	// }

	user := Human{
		lastName:  "Ergashev",
		age:       23,
		firstName: "Asadbek",
		height:    171,
		width:     82,
		birthday:  "2000-05-29",
	}

	user.age = 24

	return user
}

func makeUsers() []Human {

	var users []Human
	for i := 0; i < 4; i++ {
		users = append(users, Human{
			lastName:  "Ergashev",
			age:       i,
			firstName: "Asadbek",
			height:    171,
			width:     82,
			birthday:  "2000-05-29",
		})
	}

	return users
}

func main() {

	// user := makeSingleUser()
	// fmt.Printf("%+v\n", user)
	// fmt.Println(user)

	// inputIndex := 2
	users := makeUsers()

	// if len(users) > inputIndex {
	// 	fmt.Println(users[inputIndex])
	// } else {
	// 	fmt.Println("error: no rows")
	// }

	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
