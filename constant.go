package main

import "fmt"

//นิยามค่าคงที่
func main() {
	const name string = "Suksom" //เมื่อประกาศ name เป็นค่าคงที่ด้วย const ค่าของ name จะไม่สามารถเปลี่ยนแปลงได้
	// name = "Studio" // เมื่อเราเขียนให้ค่า name มีการเปลี่ยนแปลงจะไม่สามารถ run ได้
	fmt.Println("My Name is ", name)
}
