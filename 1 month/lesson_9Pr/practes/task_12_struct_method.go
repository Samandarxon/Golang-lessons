package main

import "fmt"

type Student struct{
	Name string
	Grades []int
	Age int
}
func (this Student)getAverageGrade()int{
	summ:=0
	for _,el:=range this.Grades{
		summ+=el
	}
	return summ/len(this.Grades)
}

func (this Student)getMaxGrade()int{
	max:=this.Grades[0]
	for _,el:=range this.Grades{
		if max<el{
			max=el
		}
	}
	return max
}

func main(){
	student := Student{
		Name:"Ikrom",
		Grades:[]int{30,45,78,65,42,12,46},
		Age:23,
	}

	fmt.Println(student)
	fmt.Println("O'rtacha qiymat :=",student.getAverageGrade())
	fmt.Println("Max qiymat :=",student.getMaxGrade())


}