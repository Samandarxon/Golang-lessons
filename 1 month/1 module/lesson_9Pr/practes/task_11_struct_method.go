package main

import "fmt"

type MyHashSet struct{
	obj []int
}

func (this MyHashSet) Add(value int )MyHashSet{
	// fmt.Println("salom")
	this.obj  = append(this.obj,value)
	return this
}


func (this MyHashSet) Contains(value int )bool{
	for _,el:=range this.obj{
		if value == el{
			return true
		}
	}
	return false
}

func (this MyHashSet) Remove(index int )MyHashSet{
		this.obj = append(this.obj[:index],this.obj[index+1:]...)	
	return this
}


func main(){
	var myHashSet MyHashSet

	myHashSet = myHashSet.Add(2)
	myHashSet = myHashSet.Add(3)
	myHashSet = myHashSet.Add(4)
	myHashSet = myHashSet.Add(5)
	myHashSet = myHashSet.Add(6)
	myHashSet = myHashSet.Add(27)
	myHashSet = myHashSet.Add(37)
	myHashSet = myHashSet.Add(47)
	fmt.Println(myHashSet)
	fmt.Println(myHashSet.Contains(4))
	myHashSet = myHashSet.Remove(3)
	fmt.Println(myHashSet)

}