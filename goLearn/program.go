package main

import "fmt"

//นิยามตัวแปร
func main() {
	var num = []int{100, 200, 300}
	name := []string{"A", "B", "C", "D", "E", "F", "G"}
	name = append(name, "C", "d", "f") //เพิ่มค่าเข้า array name โดยเพิ่ม C,d,f
	num = append(num, 400, 500, 600)

	fmt.Println(num[2:5]) //เริ่ม index 2 และจบที่ index4
	fmt.Println(name[1:])

	// array ต้องใช้เลข index ซึ่งจำยาก
	//ใช้ key เพื่อหาข้อมูลใน array ได้ง่ายขึ้น เพื่อเข้าถึงข้อมูล สามารถกำหนดได้อิสระ

	/* country := map[string]string{}
	country["TH"] = "Thailand"
	country["EN"] = "England" */

	/* coin := map[string]string{}
	coin["ETH"] = "Etherium"
	coin["BTC"] = "Bitcoin" */

	people := map[string]int{}
	people["Thailand"] = 70
	people["England"] = 50

	//coin := map[string]string{"ETH": "Ether", "BTC": "Bitcoin"}

	country := map[string]string{"TH": "Thailand", "JP": "Japan"}
	value, check := country["EN"]
	//value คือค่าที่อยู่ใน key
	//check เป็น bool

	//เช็คว่าเจอ key ไหม
	if check {
		fmt.Println(value)
	} else {
		fmt.Println("ไม่พบข้อมูล")
	}

	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Hello Go Programming", i)

	}
	fmt.Println("จบ")

	/* for index := 0; index < len(name); index++ {
		fmt.Println("Index = ", index, "value = ", name[index])
	} */

	for _, value := range name { //_ คือการ ignore มองข้ามตัวแปรนั้นไป ในที่นี้เราจะมองข้ามตัวแปร index
		println(value)
	}

	language := map[string]string{"TH": "Thai", "EN": "English"}
	for key, value := range language {
		fmt.Println("Key= ", key, "Value= ", value)
	}

	//call function
	showMessage("tea")
	showMessage("kong")
	showMessage("jojo")
	deliver := getDelivery()
	total(50, 100) //รับค่าอย่างเดียว ไม่ต้องสร้างตัวแปร
	fmt.Println("ค่าจัดส่ง ", deliver)
	myCart := getTotalCart(500, 500) //ถ้า func มีการ Return ต้องสร้างตัวแปรขึ้นมารับค่า
	fmt.Println("ยอดรวมสินค้า ", myCart)

	result, checked := summation(5, 200)
	fmt.Println(result, checked)

}

func showMessage(name string) {
	fmt.Println("Hello ", name)
}
func total(x, y int) { //func รับค่า
	//argument ค่าที่ส่งเข้าฟังก์ชัน 10,5 เป็น argument ที่ส่งเข้ามาใน parameter x,y
	fmt.Println("ยอดรวม ", x+y)
}
func getDelivery() int { //func มีการ return
	return 50
}
func getTotalCart(x, y int) int { //func รับค่าและ return
	total := x + y
	return total
}
func summation(x, y int) (int, string) {
	total := x + y
	status := ""

	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}

	return total, status
}
