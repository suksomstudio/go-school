package main

import "fmt"

/*
Slices มีลักษณะคล้าย Array แต่มีความสามารถในการปรับเปลี่ยนขนาดสมาชิกได้ (Dynamic Size)
*/

func main() {
	//การนิยาม Slices จะไม่มีการระบุขนาดจำนวนของชุดข้อมูล
	names := []string{"สมชาย", "สมหญิง", "สมจิต", "สมใจ"}
	fmt.Println("สมาชิกทั้งหมดมี = ", names)
	names = append(names, "สมปอง")  //append = การเพิ่มสมาชิกของชุดข้อมูล
	names = append(names, "สมหน้า") //append = การเพิ่มสมาชิกของชุดข้อมูล
	fmt.Println("สมาชิกทั้งหมดมี = ", names)
}
