package main

import "fmt"

// scanf การรับค่าจากคีย์บอร์ด
func main() {
	var name string

	fmt.Println("ป้อนชื่อนักเรียน = ")
	fmt.Scanf("%s", &name)

	fmt.Println("สวัสดี =", name)

}
