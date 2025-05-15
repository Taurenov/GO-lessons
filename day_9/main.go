package main

import "fmt"

func main(){
	numbers := [5]int{10,20,30,40,50}
	index := 1
	for i := 0; i < len(numbers); i++{
		fmt.Println("Элемент с индексом:", i, ":", numbers[i])
	}
	days := [7]string{"Понедельник","Вторник","Среда","Четверг","Пятница","Суббота","Воскресенье"}
	if index >= 0 && index < len(days){
		fmt.Println("День недели:", days[index])
	} else {fmt.Println("Некоректный индекс")}
	float := [10]float64{1.5,2.0,2.5,3.0,3.5,4.0,4.5,5.0,5.5,6.0}	
	var sum float64 = 0 // Инцилизация переменной для хранения суммы
	//Сложение
	for _, number := range float {
		sum += number // Добавление текущего элемента к сумме
	}
	average := sum / float64(len(float)) // Высчитываем среднее значение
	fmt.Println("Среднее значение:", average)
	num1 := [3]int{3, 7, 9}
	num2 := [3]int{10, 11, 21}
	sumArray := [3] int{} // 3 массив для хранения сумм
	for x := 0; x< len(num1); x++ {
		sumArray[x] = num1[x] + num2[x]
	}
	fmt.Println("Сумма массивов:", sumArray)
}