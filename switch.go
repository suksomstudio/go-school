package main

import "fmt"

/*
: เลือกใช้บริการระบบธนาคาร
ให้ป้อนหมายเลขตัวเลือกเพื่อใช้บริการ
- ถ้าป้อน 1 คือ การเปิดบัญชีธนาคาร
- ถ้าป้อน 2 คือ การถอน-ฝากเงิน
- ถ้าพิมพ์ตัวเลขอื่น แจ้งว่าข้อมูลไม่ถูกต้อง
*/

func main() {
	//รับ Input
	var number int

	fmt.Println("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("เปิดบัญชีใหม่")
	case 2:
		fmt.Println("ฝากเงิน - ถอนเงิน")
	default:
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}

}
