package main

import "fmt"

/*
Задание 2
Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Printf("Arr: %v\n", arr)

}