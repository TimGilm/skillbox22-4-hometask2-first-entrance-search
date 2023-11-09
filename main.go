/*Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив.
Сложность алгоритма должна быть минимальная. Что оценивается Верность индекса.
При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.*/

package main

import "fmt"

const n = 12

var a [n]int

func main() {
	fmt.Println("Введите 12 чисел в массив") //1 2 2 2 3 4 5 6 7 8 9 10
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)
		a[i] = num
	}
	fmt.Print("Введите число, индекс первого вхождения которого необходимо найти")
	var value int
	fmt.Scanln(&value)
	fmt.Printf("Индекс первого вхождения указанного числа %v", findIndex(a, value))
}

func findIndex(a [n]int, value int) (index int) {
	//сперва применим бинарный поиск, зная что наш массив отсортирован
	index = -1
	min := 0
	max := n - 1
	for max >= min {
		middle := (max + min) / 2
		if a[middle] == value {
			index = middle
			break
		} else if a[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	//проверим массив "слева" на дубли, и переназначим index, т.к. нам нужно первое вхождение элемента.
	for i := index; i >= 0; i-- {
		if a[i] < value {
			break
		} else {
			index = i
		}
	}
	return
}
