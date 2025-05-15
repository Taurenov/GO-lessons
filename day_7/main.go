package main

import "fmt"

func main(){
	//Вводные значения
	time := 22
	month := 8
	score := 88

	//1 задание время через switch
	switch {
	case time >= 6 && time <= 11 :
		fmt.Println("Доброе утро")
	case time >= 12 && time <= 17 :
		fmt.Println("Добрый день")
	case time >= 18 && time <= 22 :
		fmt.Println("Добрый вечер")
	default:
		fmt.Println("Доброй ночи")
	}
	//2 определение месяца
	switch month {
	case 1 :
		fmt.Println("Январь")
	case 2 :
		fmt.Println("Февраль")
	case 3 :
		fmt.Println("Март")
	case 4 :
		fmt.Println("Апрель")
	case 5 :
		fmt.Println("Май")
	case 6 :
		fmt.Println("Июнь")
	case 7 :
		fmt.Println("Июль")
	case 8 :
		fmt.Println("Август")
	case 9 :
		fmt.Println("Сентябрь")
	case 10 :
		fmt.Println("Октябрь")
	case 11 :
		fmt.Println("Ноябрь")
	case 12 :
		fmt.Println("Декабрь")
	default:
		fmt.Println("Некоректный ввод месяца")
	}
	//3 Задание оценивание
	switch {
	case score >=90 :
		fmt.Println("Ваша оценка A")
	case score >=80 :
		fmt.Println("Ваша оценка B")
	case score >=70 :
		fmt.Println("Ваша оценка C")
	case score >=60 :
		fmt.Println("Ваша оценка D")
	default:
		fmt.Println("Ваша оценка F")
	}
}