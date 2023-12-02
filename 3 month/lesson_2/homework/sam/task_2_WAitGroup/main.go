package main

import (
	"app/pkg/helpers"
	"fmt"
	"sync"
	"time"
)

func main() {
	x_time := time.Now()
	fetch := helpers.Fetch{
		Url: "https://cbu.uz/uz/arkhiv-kursov-valyut/json/all/2023-12-25/",
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fetch.GetRequest(&wg)
	fmt.Println("salom")
	wg.Wait()
	wg.Add(1)
	go helpers.SaveToExcel(&wg, fetch.Obj.Sorted().Objs, "sam")
	wg.Wait()
	fmt.Println(time.Since(x_time))

}
