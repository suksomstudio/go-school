package main

import "fmt"

/*
Array
- เป็นชุดคำสั่งตัวแปรที่อยู่ในรูปลำดับใช้เก็บค่าข้อมูลให้อยู่ในกลุ่มเดียวกัน โดยข้อมูลภายใน Array จะถูกเก็บในตำแหน่งที่ต่อเนื่องกัน
- เป็นตัวแปรที่ใช้ในการเก็บข้อมูลที่มีลำดับต่อเนื่อง ซึ่งข้อมูลมีค่าได้ หลายค่าโดยใช้อ้างอิง ได้เพียง ชื่อเดียว และใช้หมายเลขกำกับ (index)ให้กับตัวแปรเพื่อ จำแนกความแตกต่างของค่าตัวแปรแต่ละตัว
*/

func main() {
	//สรา้ง Array แบบยาว
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)
	sizenumber := len(numbers) //len บอกจำนวนขนาดของ Array
	fmt.Println("จำนวนตัวเลขทั้งหมด = ", sizenumber)

	//สรา้ง Array แบบสั้น
	pets := [4]string{"หมู", "หมา", "กา", "ไก่"}
	fmt.Println(pets)
	sizepets := len(pets) //len บอกจำนวนขนาดของ Array
	fmt.Println("จำนวนสัตย์เลี้ยงทั้งหมด = ", sizepets)

	fmt.Println(numbers[0]) //ระบุ index ที่ต้องการใช้งาน
	fmt.Println(numbers[3]) //ระบุ index ที่ต้องการใช้งาน
	fmt.Println(pets[0])    //ระบุ index ที่ต้องการใช้งาน
	fmt.Println(pets[3])    //ระบุ index ที่ต้องการใช้งาน
}
