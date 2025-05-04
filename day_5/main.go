package main

import "fmt"

func main() {
	num1 := 55
	num2 := 33
	floatNum1 := 3.3
	floatNum2 := 1.1

	fmt.Println("--- Арефмитические Операторы ---")
	fmt.Println("Сумма", num1+num2)
	fmt.Println("Сумма дроби", floatNum1+floatNum2)
	fmt.Println("Разность", num1-num2)
	fmt.Println("Разность дроби", floatNum1-floatNum2)
	fmt.Println("Умножение", num1*num2)
	fmt.Println("Умножение дроби", floatNum1*floatNum2)
	fmt.Println("Деление", num1/num2)
	fmt.Println("Деление дроби", floatNum1/floatNum2)
	fmt.Println("Остаток", num1%num2)

	increment := 5
	fmt.Println("Инкримент(постфиксный):", increment)
	increment++
	fmt.Println("Значение increment после постфикса:", increment)
	
	decrement := 8
	fmt.Println("Декремент(Префиксный):", decrement)
	decrement--
	fmt.Println("Декремент постфикса:", decrement)

	//Сравнение

	a := 5
	b := 7
	str1 := "hello"
	str2 := "World"
	str3 := "hello"
	fmt.Println("Равно?:", a == b)
	fmt.Println("Не равно?:", a != b)
	fmt.Println("А > B", a > b)
	fmt.Println("A < B", a < b)
	fmt.Println("A >= B", a >= b)
	fmt.Println("A <= B", a <= b)
	fmt.Println("Слова одинаковы?", str1 == str3)
	fmt.Println("Слова не одинаковые?", str1 != str2)

	p := true
	q := false

	fmt.Println("p and q", p && q)
	fmt.Println("p or q", p || q)
	fmt.Println("not p", !p)
	fmt.Println("not (p and q)", !(p && q))
}