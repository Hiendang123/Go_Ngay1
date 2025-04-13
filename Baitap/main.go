package main

import (
	"fmt"
	"math"
)

func calculateArea(r float64) float64 {
	fmt.Printf("Nhập bán kính: ")
	fmt.Scan(&r)
	return math.Pi * r * r
}

func main() {
	var r float64
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("Số chẵn là: %d\n", i)
		} else {
			fmt.Printf("Số lẻ là: %d\n", i)
		}
	}

	fmt.Println("Diện tích hình tròn là: " + fmt.Sprintf("%.2f", calculateArea(r)))

	var ageCustomer int

	fmt.Print("Nhập tuổi của khách hàng: ")

	fmt.Scanf("%d", &ageCustomer)

	switch {
	case ageCustomer < 5:
		fmt.Println("vé miễn phí")
	case ageCustomer > 5 && ageCustomer < 22:
		fmt.Println("vé $10")
	case ageCustomer > 22:
		fmt.Println("vé $15")
	default:
		fmt.Println("không có vé")
	}

	var toan, van, anh, tbt float64
	fmt.Println("Nhập điểm Toán:")
	fmt.Scan(&toan)
	fmt.Println("Nhập điểm Văn:")
	fmt.Scan(&van)
	fmt.Println("Nhập điểm Anh:")
	fmt.Scan(&anh)

	tbt = (toan + van + anh) / 3

	switch {
	case tbt >= 9:
		fmt.Println("Xuất sắc")
	case tbt >= 7 && tbt < 9:
		fmt.Println("Giỏi")
	case tbt >= 5 && tbt < 7:
		fmt.Println("Khá")
	case tbt < 5:
		fmt.Println("Trung bình")
	default:
		fmt.Println("Không có điểm")
	}

	var num int

	fmt.Print("Nhập một số nguyên để kiểm tra: ")
	fmt.Scan(&num)

	if num < 2 {
		fmt.Println("Not Prime")
		return
	}

	isPrime := true
	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			isPrime = false
		}
	}

	if isPrime {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}
