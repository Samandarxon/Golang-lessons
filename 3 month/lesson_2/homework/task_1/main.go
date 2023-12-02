package main

import (
	"app/pkg/helpers"
	"fmt"
	"time"
)

func main() {
	x_time := time.Now()
	fetch := helpers.Fetch{
		Url: "https://cbu.uz/uz/arkhiv-kursov-valyut/json/all/2023-12-25/",
	}
	fetch.GetRequest()
	helpers.SaveToExcel(fetch.Obj.Sorted().Objs, "sam")
	fmt.Println(time.Since(x_time))
}
