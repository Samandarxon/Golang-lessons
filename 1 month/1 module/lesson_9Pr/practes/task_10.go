package main

import "fmt"

func repeatedCharacter(s string) byte {
	var s_by byte
	var x_map = make(map[string]int)
	fmt.Println(s_by)
	for i, el := range s {
		fmt.Println(i, i+1)
		if i < len(s)-1 && string(s[i+1]) == string(el) {
			s_by = byte(el)
			break
		}else{
			x_map[string(el)]++
		}
	}
	if s_by == 0 {
		min:=int([]byte("z")[0])
		fmt.Println(min)
		for key, value := range x_map{
			if value>1 {

				// s_by=[]byte(key)[0]
				if  min>=int([]byte(key)[0]){					
					min=int([]byte(key)[0])
					s_by=[]byte(key)[0]
					// break
                }else{
					s_by=[]byte(key)[0]
					break
                    
                }
			}
		}
	}
	return s_by
}

func main() {
	var str = "regzueqr"
	fmt.Printf("%c\n", repeatedCharacter(str))
}
