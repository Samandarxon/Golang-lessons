package main

import "fmt"

func main() {
	// tastk_1(3)
	// tastk_1(5)
	// tastk_1(15)
	// tastk_1(13)
	// tastk_2(123)
	 task_3(2, 5, 7)
	// task_4(3, 3, 4)
	// task_4(23, 3, 45)
	// task_4(22, -44, -44)
	// task_5(3, 9, 4)
	// task_5(23, 3, 23)
	// task_5(22, -44, 245)
	// task_6(6)
	// task_6(7)
	// task_7(1996)
	// task_7(2024)
	// task_7(2003)
	// task_7(2002)
	// task_8(2, 4, 4)
	// task_9(2, 3, "+")
	// task_9(2, 3, "-")
	// task_9(3, 0, "/")
	// task_9(2, 3, "*")
	// task_9(2, 3, "$")
	// task_10()
	// task_11(15)
	// task_12(40)
	// task_13(5)
	task_14(324)
}

func tastk_1(n int) {
	if n%5 == 0 && n%3 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(n)
	}
}
func tastk_2(n int) {
	a, b, c := n/100, n%100/10, n%10
	d := c*100 + b*10 + a
	if d == n {
		fmt.Println("Palidrom")
	} else {
		fmt.Println("It isn't Palidrom")
	}
}

func task_3(a, b, c int) {
	fmt.Printf("a =: %v b =: %v c =: %v\n", a, b, c)
	a, b, c = c, a, b
	fmt.Printf("a =: %v b =: %v c =: %v\n", a, b, c)
}

func task_4(x, y, z int) {
	if x <= y && x <= z {
		fmt.Println(x)
	} else if y <= x && y <= z {
		fmt.Println(y)
	} else if z <= x && z <= y {
		fmt.Println(z)
	}

}
func task_5(x, y, z int) {
	if x >= y && x >= z {
		fmt.Println(x)
	} else if y >= x && y >= z {
		fmt.Println(y)
	} else if z >= x && z >= y {
		fmt.Println(z)
	}

}
func task_6(n int) {
	if n%2 == 0 {
		fmt.Println("even \"Juft\"")
	} else {
		fmt.Println("odd \"Toq\"")
	}
}
func task_7(year int) {
	if year%4 == 0 {
		fmt.Println("Kabissa yil")
	} else {
		fmt.Println("Kabissa yil emas")
	}
}

func task_8(A, B, C float32) {
	fmt.Println((A + B + C) / 3)
}

func task_9(a, b float32, operator string) {
	if operator == "+" {
		fmt.Printf("%v + %v = %v\n", a, b, a+b)
	} else if operator == "-" {
		fmt.Printf("%v - %v = %v\n", a, b, a-b)
	} else if operator == "*" {
		fmt.Printf("%v * %v = %v\n", a, b, a*b)
	} else if operator == "/" && b != 0 {
		fmt.Printf("%v / %v = %v\n", a, b, a/b)
	} else if operator == "/" && b == 0 {
		fmt.Printf("%v sonini %v soniga Bo'lib bo'lmaydi\n", a, b)
	} else {
		fmt.Printf("\"%s\" Bunaqa matematik operator yo'q\n", operator)
	}
}
func task_10() {
	var (
		investorPercentage = 70
		companyPercentage  = 30

		malibuPrice   = 700_000
		insuranePrice = 100_000
		orderDay      = 2

		totalSum  = malibuPrice*orderDay + insuranePrice*orderDay
		investor  = totalSum * investorPercentage / 100
		company   = totalSum*companyPercentage/100 - insuranePrice*orderDay
		insurance = insuranePrice * orderDay
	)

	fmt.Println("Total Summa: ", totalSum)
	fmt.Println("Investor: ", investor)
	fmt.Println("Company: ", company)
	fmt.Println("Insurance: ", insurance)
}

func task_11(a int) {
	fmt.Println("S = ", a*a, "\nP = ", 4*a)
}

func task_12(selse float32) {
	fmt.Println("Kelven =", selse*273.15)
	fmt.Println("Fahrenheit = ", selse*1.8+32)
}

func task_13(day int) {
	if day == 1 {
		fmt.Println("Dushanba")
	} else if day == 2 {
		fmt.Println("Seshanba")
	} else if day == 3 {
		fmt.Println("Chorshanba")
	} else if day == 4 {
		fmt.Println("Payshanba")
	} else if day == 5 {
		fmt.Println("Juma")
	} else if day == 6 {
		fmt.Println("Shanba")
	} else if day == 7 {
		fmt.Println("Yakshanba")
	}
}

func task_14(n int) {
	fmt.Println(n)
	a, b := n%100/10, n%10
	a, b = b, a
	fmt.Println(n - n%100 + a*10 + b)
}
