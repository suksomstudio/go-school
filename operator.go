package main

import "fmt"

// operator ตัวดำเนินการทางคณิตศาสตร์
func main() {
	var number1 int = 10
	var number2 int = 3

	fmt.Println("ผลบวก = ", number1+number2)
	fmt.Println("ผลลบ = ", number1-number2)
	fmt.Println("ผลคูณ = ", number1*number2)
	fmt.Println("ผลหาร = ", number1/number2)
	fmt.Println("เศษ = ", number1%number2)

}
