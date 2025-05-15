package main

import "fmt"

func main(){
	// 1 задание //
	number := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Объявил Массив
	slice1 := number[0:5] // Делаем 1 слайс и задаём его
	slice2 := number[5:10] // Делаем 2 слайс и задаём его

	fmt.Println("Первые 5 элементов:", slice1)
	fmt.Println("Длина slice1:", len(slice1))
	fmt.Println("Ёмкость slice1:", cap(slice1))

	fmt.Println("Вторые 5 элементов:", slice2)
	fmt.Println("Длина slice2:", len(slice2))
	fmt.Println("Ёмкость slice2:", cap(slice2))

	// 2 задание //
	var slice3 []int // Создаю пустой слайс целых чисел
	// Цикл for для чисел от 1 до 5
	for i := 1; i <= 5; i++ {
		square := i*i // Вычесдяем квадрат числа
		slice3 = append(slice3, square) // Добовляем квадрат в слайс
	}
	fmt.Println("Слайс квадратов:", slice3)

	// 3 задание //
	slice4 := []string{"один", "два", "три"} // Создаем слайс строк
		fmt.Println("Начальная длина:", len(slice4))
		fmt.Println("Начальная ёмкость:", cap(slice4))

	slice4 = append(slice4, "четыре", "пять") // Добавляем 2 элемента
		fmt.Println("Созданный слайс:", slice4)
		fmt.Println("Конечная длина:", len(slice4))
		fmt.Println("Конечная ёмкость:", cap(slice4))
		
	// 4 задание //
	sourceSlice := []int{1, 2, 3}
	destinationSlice := make([]int, len(sourceSlice)) // Создаем целевой слайс той же длины
	numCopied := copy(destinationSlice, sourceSlice)

	fmt.Println("Исходный слайс:", sourceSlice)
	fmt.Println("Скопированный слайс:", destinationSlice)
	fmt.Println("Количество скопированных элементов:", numCopied)

	// Пример копирования в слайс меньшей длины:
	destinationSliceShort := make([]int, 2)
	numCopiedShort := copy(destinationSliceShort, sourceSlice)
	fmt.Println("Короткий целевой слайс:", destinationSliceShort)
	fmt.Println("Скопировано в короткий:", numCopiedShort)
}