package main

import "fmt"

func main(){
	//Водные данные:
num := 11
time := 22
year := 2024
score := 79

	//1 блок задачи число равно или нет 0
	if num > 0 {
		fmt.Println("Число положительное")
	} else if num < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно 0")
	}
	//2 Блок время и приветсвие от его числа
	if time >= 6 && time <=11 {
		fmt.Println("Доброе утро")
	} else if time >= 12 && time <= 17 {
		fmt.Println("Добрый день")
	} else if time >= 18 && time <= 22 {
		fmt.Println("Добрый вечер")
	} else {
		fmt.Println("Доброй ночи")
	}
	//3 часть весокосный год
	if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
		fmt.Println("Год весокосный")
	} else {
		fmt.Println("Год не весокосный")
	}
	//4 Ставим оценку по баллам
	if score >=90 {
		fmt.Println("Твоя оценка A")
	} else if score >=80 {
		fmt.Println("Тво оценка B")
	} else if score >=70 {
		fmt.Println("Тво оценка C")
	} else if score >= 60 {
		fmt.Println("Тво оценка D")
	} else {
		fmt.Println("Тво оценка F")
	}
}