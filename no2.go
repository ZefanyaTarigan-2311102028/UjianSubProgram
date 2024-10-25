// 2311102028
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan jumlah kelopak (a): ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan jumlah kelopak (b): ")
	fmt.Scanln(&b)

	fmt.Print("Bunga Sempurna antara ", a, " dan ", b, ": ")
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isPerfectNumber(num int) bool {
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
		}
	}
	return sum == num
}
