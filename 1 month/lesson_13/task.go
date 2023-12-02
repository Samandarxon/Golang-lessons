package main

import "fmt"

func direct(x interface{})(interface{},bool){
	key:=""
	if _,ok := x.(string);ok {
		key=x.(string)
	}
	x_key:=
	if _,ok := x.(int);ok {
		x_key=x.(int)
	}
	if key==""{
		switch key{
		case "TOP","0":
			return 0,false
		case "BOTTOM","1":
			return 1,false
		case "LEFT","2":
			return 2,false
		case "RIGHT","3":
			return 3,false
		default:
			return fmt.Sprint("\"\""),true
	} 

	}
	switch key{
	case "TOP","0":
		return 0,false
	case "BOTTOM","1":
		return 1,false
	case "LEFT","2":
		return 2,false
	case "RIGHT","3":
		return 3,false
	default:
		return fmt.Sprint("\"\""),true
} 
}
func main() {

  // 0 TOP
  // 1 BOTTOM
  // 2 LEFT
  // 3 RIGHT

  result, err := direct("BOTTOM")
  fmt.Println(result, err) // 1 false

  result, err = direct(3)
  fmt.Println(result, err) // RIGHT false

  result, err = direct(true)
  fmt.Println(result, err) // "" true
}