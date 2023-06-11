package main

import "fmt"

/*
Map
ตัวแปรที่เก็บข้อมูลในรูปแบบคู่ key, value มีลักษณะคล้าย Array แต่จะใช้ key เป็น index เพื่อเชื่อมโยงข้อมูล (value) ที่เก็บใน key นั้นๆ ถ้าทราบ key ก็สามารถเข้าถึง value หรือข้อมูลได้นั่นเอง
*/

func main() {
	// การนิยาม  Map
	country := map[string]string{"th": "ประเทศไทย", "jp": "ประเทศญี่ปุ่น"} // th, jp => KEY และ ประเทศไทย, ประเทศญี่ปุ่น => Value
	fmt.Println(country["jp"])
	fmt.Println(country["th"])
	fmt.Println(country)

	value, check := country["EN"]

	if check {
		fmt.Println(value)
	} else {
		fmt.Println("ไม่พบข้อมูล")
	}
}
