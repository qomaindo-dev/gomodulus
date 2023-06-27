package main

import (
	"fmt"
)

func main() {
	var pja, k1, k2 int

	fmt.Println("Input Panjang Angka: ")
	_, err := fmt.Scanf("%d", &pja)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pja <= 0 {
		fmt.Println("Panjang angka harus lebih dari nol (0)")
		return
	}

	fmt.Println("Input Kelipatan Ke-1: ")
	_, errK1 := fmt.Scanf("%d", &k1)
	if errK1 != nil {
		fmt.Println(errK1)
		return
	}

	if k1 <= 0 {
		fmt.Println("Kelipatan tidak boleh kosong")
		return
	}

	fmt.Println("Input Kelipatan Ke-2: ")
	_, errK2 := fmt.Scanf("%d", &k2)
	if errK2 != nil {
		fmt.Println(errK2)
		return
	}

	if k2 <= 0 {
		fmt.Println("Kelipatan tidak boleh kosong")
		return
	}

	fmt.Println("******Hasil******")
	for i := 1; i <= pja; i++ {
		if ((i % k1) == 0) && ((i % k2) == 0) {
			fmt.Println("TIKI-TAKA")
		} else if (i % k1) == 0 {
			fmt.Println("TIKI")
		} else if (i % k2) == 0 {
			fmt.Println("TAKA")
		} else {
			fmt.Println(i)
		}
	}
}
