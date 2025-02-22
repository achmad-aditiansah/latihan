package logic

import (
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

	for i := 0; i <n; i++ {
		slice[i] = num
		num += 3
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola4(n int) []int {
	slice := make([]int, n)

	num := 2 * n -1

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

	for i := 0; i < n / 2; i++ {
		slice[i] = num
		num += 2
	}

	num = 2

	for i := n - 1; i >= n/2; i--{
		slice[i] = num
		num += 2
	}
	printer.PrintSlice(slice)
	return slice
}

func Logic1Pola9(n int) []int{
	slice := make([]int, n)

	num := 3

	for i := 0; i < n / 2; i++ {
		slice[i] = num
		num += 3
	}

	num = 3

	for i := n - 1; i >= n / 2; i-- {
		slice[i] = num
		num += 3
	}
	printer.PrintSlice(slice)
	return slice
}
