package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	/* circum := calCircum(7)
	fmt.Printf("circumference of the circle: %.2f\n", circum)

	circum2 := calCircum2(7)
	fmt.Printf("circumference of the circle: %.2f\n", circum2)

	fmt.Println("Input radius: ")
	var circum3 int
	fmt.Scanln(&circum3)
	calCircum3(circum3)

	fmt.Println("\n Input x, y, z ")
	var x, y, z int
	fmt.Scanln(&x, &y, &z)
	//fmt.Println(x, y, z)
	fmt.Println(checkNum(x, y, z))

	fmt.Println(powNum()) */

	/* fmt.Println("Enter size array: ")
	var size int
	fmt.Scanln(&size)
	var member = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Println("Input number: ", i+1)
		fmt.Scanln(&member[i])
	}
	fmt.Println(sumArrays2(member)) */

	calculateHalfHalf()
}
func calCircum(radius int64) float64 {

	pi := math.Pi
	circum := (2 * pi) * float64(radius)

	return circum

}

func calCircum2(radius int) float64 {
	return 2 * 3.14 * float64(radius)
}

func calCircum3(radius int) {
	fmt.Printf("%.2f", 2*3.14*float64(radius))
}

func checkNum(x, y, z int) float64 {
	if x == 0 {
		return float64(x + y)
	} else if z == 1 {
		return float64(x - y)
	} else {
		return math.Pow(float64(x), float64(y))
	}
}

func powNum() float64 {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input number: ")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	pow2 := math.Pow(float64(x), 2)
	fmt.Printf("%d ^ 2= %.1f\n", x, pow2)
	return pow2
}

func sumArrays() {

	fmt.Println("Enter size array: ")
	var size int
	fmt.Scanln(&size)
	var member = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Println("Input number: ", i+1)
		fmt.Scanln(&member[i])
	}

	var sum int
	for i := 0; i < len(member); i++ {
		sum += member[i]
	}

	fmt.Println(sum)
}

func sumArrays2(arr []int) int64 {

	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return int64(sum)
}

func calculateHalfHalf() []float64 {
	var totalHalfHalf = make([]float64, 2)
	var give float64

	fmt.Println("โควต้าวันนี้: ")
	fmt.Scanln(&give)

	for {
		var pay float64

		fmt.Println("ใช้จ่าย: ")
		fmt.Scanln(&pay)

		if give < 0 {
			break
		}

		if (pay / 2) <= give { //ใช้คนละครึ่งแล้วพอดีกับโควตาหรือน้อยกว่า
			totalHalfHalf[0] = pay / 2
			totalHalfHalf[1] = give - pay/2
		} else {
			totalHalfHalf[0] = pay - give
			totalHalfHalf[1] = 0 //ยังไงก็เหลือ 0 เพราะใช้เกิน
		}
		fmt.Printf("จ่ายเอง %.2f เหลือโควตาวงเงิน %.2f \n", totalHalfHalf[0], totalHalfHalf[1])
	}

	return totalHalfHalf
}
