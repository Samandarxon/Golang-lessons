package main

import "fmt"

type Engineer struct {
	Name string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}
func (this Engineer)Print(){
	fmt.Printf("%+v\n",this)
}
func (this Engineer)GetProjectPriority(){

	fmt.Printf("%+v\n",this.Project.Priority)
}
func main() {
	engineer := Engineer{
		Name:"Mirzohid",
		Age:23,
		Project:Project{
			Name:"Onlayin taksi app yaratish",
			Priority:"Vaqtni tejash",
			Technologies:[]string{
				"Golang",
				"Flutter",
				"Gin",
				"Doker",
				"Microserves",
			},
		},
	}

	engineer.Print()
	engineer.GetProjectPriority()

}
