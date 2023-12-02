package main

import "fmt"

type User struct {
	name     string
	age      int
	birthday string
}
type AllUserData struct{
	users []User
}
func (d AllUserData) generateUser(limit int)[]User{
	var users = d.users
	var start,year =1,2023
	age:=start
	for i:=1; i<=limit; i++{
		if age>100{
			age = start
		}
		users=append(users,User{
			name:fmt.Sprintf("Malik-%d",i),
			age: age,
			birthday:fmt.Sprintf("%d",year-age),
		})
		age++
	}
	return users
}

func (d AllUserData)useAgeFilter(fromAge,toAge int)[]User{
	var user_slice []User

	for _,user :=range d.users{
		if user.age>=fromAge && user.age<=toAge{
			user_slice = append(user_slice,user)
		}
	}

	return user_slice
}
func main() {
	var AllData AllUserData 
	AllData.users= AllData.generateUser(100)
	// for _,user:=range AllData.users{
	// 	fmt.Printf("%+v\n",user)
	// }
	for _,user:=range AllData.useAgeFilter(10,20){

		fmt.Printf("%+v\n",user)
	}
}
