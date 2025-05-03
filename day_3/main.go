package main

import "fmt"

func main() {
	// Объявление переменных с явным указанием типа
	var integerVarExplicit int = 10
	var integer64VarExplicit int64 = 10000000000
	var float32VarExplicit float32 = 3.14
	var float64VarExplicit float64 = 2.71828
	var booleanVarExplicit bool = true
	var stringVarExplicit string = "Hello Go!"
	
	//объявление коротких переменных
	age := 23
	pi := 3.14
	hungr := false
	name := Mikhail

	// Объявление переменных с краткой записью (type inference)
	integerVarInferred := 20
	floatVarInferred := 1.618
	booleanVarInferred := false
	stringVarInferred := "World"

	// Вывод значений
	fmt.Println("--- Объявленные явно ---")
	fmt.Println("integerVarExplicit:", integerVarExplicit)
	fmt.Println("integer64VarExplicit:", integer64VarExplicit)
	fmt.Println("float32VarExplicit:", float32VarExplicit)
	fmt.Println("float64VarExplicit:", float64VarExplicit)
	fmt.Println("booleanVarExplicit:", booleanVarExplicit)
	fmt.Println("stringVarExplicit:", stringVarExplicit)

	fmt.Println("\n--- Объявленные с type inference ---")
	fmt.Println("integerVarInferred:", integerVarInferred)
	fmt.Println("floatVarInferred:", floatVarInferred)
	fmt.Println("booleanVarInferred:", booleanVarInferred)
	fmt.Println("stringVarInferred:", stringVarInferred)
	
	fmt.Println("\n--- Объявление Задание №2 с краткой записью ---")
	fmt.Println("Возраст", age)
	fmt.Println("Число пи", pi)
	fmt.Println("Голод", hungr)
	fmt.Println("Имя", name)

	// Арифметические операции с целыми числами
	num1 := 15
	num2 := 7
	fmt.Println("\n--- Арифметические операции с целыми числами ---")
	fmt.Println("Сумма:", num1+num2)
	fmt.Println("Разность:", num1-num2)
	fmt.Println("Умножение:", num1*num2)
	fmt.Println("Деление:", num1/num2)

	// Арифметические операции с числами с плавающей точкой
	floatNum1 := 3.5
	floatNum2 := 1.5
	fmt.Println("\n--- Арифметические операции с числами с плавающей точкой ---")
	fmt.Println("Сумма:", floatNum1+floatNum2)
	fmt.Println("Разность:", floatNum1-floatNum2)
	fmt.Println("Умножение:", floatNum1*floatNum2)
	fmt.Println("Деление:", floatNum1/floatNum2)

	// Булевы операции
	boolVal1 := true
	boolVal2 := false
	fmt.Println("\n--- Булевы операции ---")
	fmt.Println("AND:", boolVal1 && boolVal2)
	fmt.Println("OR:", boolVal1 || boolVal2)
	fmt.Println("NOT boolVal1:", !boolVal1)
	fmt.Println("NOT boolVal2:", !boolVal2)

	// Операции со строками
	str1 := "Hello"
	str2 := "Go"
	fmt.Println("\n--- Операции со строками ---")
	fmt.Println("Конкатенация:", str1+ " " + str2)
	fmt.Println("Длина str1:", len(str1))
	fmt.Println("Длина str2:", len(str2))

	//Арефмтика задание №2
	numb1 := 10
	numb2 := 5
	floatNumb1 := 15.5
	floatNumb2 := 22.2
	fmt.Println("Сложение", numb1+numb2)
	fmt.Println("Вычетание", numb1-numb2)
	fmt.Println("Умножение", numb1*numb2)
	fmt.Println("Деление", numb1/numb2)
	fmt.Println("Сложение", floatNumb1+floatNumb2)
	fmt.Println("Вычетание", floatNumb1-floatNumb2)
	fmt.Println("Умножение", floatNumb1*floatNumb2)
	fmt.Println("Деление", floatNumb1/floatNumb2)

	//Логика с краткой записью
	isSunny := true
	isWarm := false
	fmt.Println("\n--- Булевы операции ---")
	fmt.Println("AND:", isSunny && isWarm)
	fmt.Println("OR:", isSunny || isWarm)
	fmt.Println("NOT isSunny:", !isSunny)
	fmt.Println("NOT isWarm:", !isWarm)

	//Строки с краткой записью
	
}