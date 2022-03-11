package main

import (
	"fmt"
	"strconv"
)

func main() {
	hasil := []string{}
	var angka int
	const div3 int = 3
	const div5 int = 5
	fmt.Print("Input Angka : ")
	fmt.Scanf("%d", &angka)
	for i := 1; i <= angka; i++ {
		if i % div3 == 0 && i % div5 == 0 {
			hasil = append(hasil, "Kulina x Food")
		} else if i % div3 == 0 {
			hasil = append(hasil, "Kulina")
		} else if i % div5 == 0 {
			hasil = append(hasil, "Food")
		} else {
			hasil = append(hasil, strconv.Itoa(i))
		}
	}
	fmt.Printf("%+q", hasil)

}
