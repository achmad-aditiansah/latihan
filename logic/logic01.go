package logic

import (
	"fmt"
	printer "github.com/achmad-aditiansah/slice-printer"
)

func Logic1Pola1(n int) []int {
	slice := make([]int, n)

	num := 1

	for i := 0; i < n; i++ {
		slice[i] = num
		num += 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola2(n int) []int {
	slice := make([]int, n)

	num := 2

	for i := 0; i < n; i++ {
		slice[i] = num
		num += 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola3(n int) []int {
	slice := make([]int, n)

	num := 3

	for i := 0; i < n; i++ {
		slice[i] = num
		num += 3
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola4(n int) []int {
	slice := make([]int, n)

	num := 2*n - 1

	for i := 0; i < n; i++ {
		slice[i] = num
		num -= 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola5(n int) []int {
	slice := make([]int, n)

	num := n * 2

	for i := 0; i < n; i++ {
		slice[i] = num
		num -= 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola6(n int) []int {
	slice := make([]int, n)

	num := n * 3

	for i := 0; i < n; i++ {
		slice[i] = num
		num -= 3
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola7(n int) []int {
	slice := make([]int, n)

	num := 1

	for i := 0; i < n/2; i++ {
		slice[i] = num
		num += 2
	}

	num = 1

	for i := n - 1; i >= n/2; i-- {
		slice[i] = num
		num += 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola8(n int) []int {
	slice := make([]int, n)

	num := 2

	for i := 0; i < n/2; i++ {
		slice[i] = num
		num += 2
	}

	num = 2

	for i := n - 1; i >= n/2; i-- {
		slice[i] = num
		num += 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola9(n int) []int {
	slice := make([]int, n)

	num := 3

	for i := 0; i < n/2; i++ {
		slice[i] = num
		num += 3
	}

	num = 3

	for i := n - 1; i >= n/2; i-- {
		slice[i] = num
		num += 3
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola10(n int) []string {
	slice := make([]string, n)

	num := 2

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			slice[i] = fmt.Sprintf("%d", num)
			num *= 2
		} else {
			slice[i] = "fizz"
		}
	}
	fmt.Println(slice)
	return slice
}

func Logic1Pola11(n int) []string {
	slice := make([]string, n)
	num := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			slice[i] = "buzz"
		} else {
			slice[i] = fmt.Sprintf("%d", num)
			num *= 3
		}
	}
	fmt.Println(slice)
	return slice
}

func Logic1Pola12(n int) []int {
	slice := make([]int, n)
	pola := []int{1, 3, 5, 7}

	for i := 0; i < n; i++ {
		slice[i] = pola[i%4]
	}
	printer.PrintSlice(slice)
	return slice
}
