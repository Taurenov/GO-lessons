package main

import "fmt"

func main(){
	for i := 1; i <= 10; i++ {
		fmt.Println("Значение i:", i)
	}
	for y := 0; y <= 20; y++ {
		if y % 2 == 0 {
		fmt.Println("Значение y:", y)
		}
	}
	x := 33
	for z := 0; z <= x; z++ {
		fmt.Println("Значение z:", z)
	}
	q := 0
	for {
		fmt.Println("Значение q в бесконечном цикле:", q)
		q++
		if q > 7 {
			break // Выходим с цикла когда q = 7
		}
	}
	fmt.Println("Выход из бесконечного цикла")
	
	for w := 0; w <= 5; w++ {
		if w == 3 {
			fmt.Println("Пропуск 3", w)
			continue
		}
		fmt.Println("Значение w:", w)
	}
}