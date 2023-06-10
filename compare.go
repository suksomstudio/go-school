package main

import "fmt"

func main() {
	number1, number2 := 10, 3
	// ตัวดำเนินการเปรียบเทียบข้อมูล
	fmt.Println("เท่ากันหรือไม่ = ", number1 == number2)
	fmt.Println("ไม่เท่ากันหรือไม่ = ", number1 != number2)
	fmt.Println(number1, " มากกว่า ", number2, " = ", number1 > number2)
	fmt.Println(number1, " น้อยกว่า ", number2, " = ", number1 < number2)
	fmt.Println(number1, " มากกว่าเท่ากับ ", number2, " = ", number1 >= number2)
	fmt.Println(number1, " น้อยกว่าเท่ากับ ", number2, " = ", number1 <= number2)
}
