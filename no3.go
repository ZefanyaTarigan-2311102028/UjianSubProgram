//2311102029

package main

import "fmt"

func main() {
	var p, q int
	fmt.Print("Masukkan nilai p (kelipatan sesi pelatihan): ")
	fmt.Scanln(&p)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scanln(&q)
	fmt.Println("Jumlah sesi pelatihan dalam setahun:", countSessions(1, 365, p, q))
}

func countSessions(day, maxDay, p, q int) int {
	if day > maxDay {
		return 0
	}
	count := 0
	if day%p == 0 && day%q != 0 {
		count++
	}
	return count + countSessions(day+1, maxDay, p, q)
}
