package main

import "fmt"

func main() {

	//การแสดงผลชนิดข้อมูลของตัวแปร
	var name string = "Suksom"
	var age int = 25
	var score float64 = 95.8
	var isPass bool = true
	// %v = การแสดงค่าที่มีอยู่ในตัวแปร
	// %T = การแสดงชนิดของข้อมูล
	// \n = การขึ้นบันทัดใหม่ของการแสดงผลสำหรับ Printf
	fmt.Printf("My Name is %v \n", name)
	fmt.Printf("Age = %v \n", age)
	fmt.Printf("Score = %v \n", score)
	fmt.Printf("Pass Exam = %v \n", isPass)
	fmt.Printf("Type Name = %T \n", name)
	fmt.Printf("Type Age = %T \n", age)
	fmt.Printf("Type Score = %T \n", score)
	fmt.Printf("Type isPass = %T \n", isPass)
}
