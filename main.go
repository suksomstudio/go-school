package main

import "fmt" //นำเอา package เข้ามาทำงาน
/*
ให้ทำงานในฟังก์ชั่น main
มีการแสดงผลออกทางจอภาพ
โดยการแสดงผลจะอาศัยคำสั่ง fmt.println
*/

func main() {
	//การนิยามตัวแปรแบบ Manual Type Declaration
	var name string = "Suksom"
	var age int = 25
	var score float64 = 95.8
	var isPass bool = true

	fmt.Println("Hello Go Programming")
	fmt.Println("Welcome to Suksom studio")

	fmt.Println("My Name is ", name)
	fmt.Println("Age ", age)
	fmt.Println("Score ", score)
	fmt.Println("Pass Exam ", isPass)

	//การนิยามตัวแปรแบบ Type Inference
	name1 := "Studio"
	age1 := 21
	score1 := 96.8
	isPass1 := true

	fmt.Println("My Name is ", name1)
	fmt.Println("Age ", age1)
	fmt.Println("Score ", score1)
	fmt.Println("Pass Exam ", isPass1)
}
