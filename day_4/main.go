package main

import "fmt"

//3.Область видимости:
var packageName string = "main"

func main() {
	// 1.Различные способы объявления
	age := 23
	city := "SPB"
	inStudent := true
	//Вывод значений
	fmt.Println("Возраст", age)
	fmt.Println("Город", city)
	fmt.Println("Правди ли студент?", inStudent)
	//2.Множественное объявление и присваивание
	var firstName string
	var lastName string
	var zipCode int
	firstName = "Михаил"
	lastName = "Исатаев"
	zipCode = 5288
	fmt.Println("Имя", firstName)
	fmt.Println("Фамилия", lastName)
	fmt.Println("Пин-код", zipCode)

	x, y := 3.14, 2.12345
	fmt.Println("Значение X", x)
	fmt.Println("Значение Y", y)

	//3 Область видимости:
	if true {
		packageName := "innerScope"
		fmt.Println("Локальное значение внутри if", packageName)
	}
	fmt.Println("глобальное значение вне if", packageName)

	//4 Константы:
	const gravity float64 = 9.8
	const appName string = "MyGoApp"
	const maxUsers = 100
	fmt.Println("Значение константы Gravity", gravity)
	fmt.Println("Значение константы AppName", appName)
	fmt.Println("Значение константы MaxUser", maxUsers)
}