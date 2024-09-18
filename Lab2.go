package main

import (
	"fmt"
	"math"
)

func main() {
	//Задание 1
	fmt.Println("Введите целое число")
	var a int
	fmt.Scanln(&a)
	fmt.Println("Проверка на четность: ", isEven(a), "\n")

	//Задание 2
	fmt.Println("Проверка на знак: ", check(a), "\n")

	//Задание 3
	fmt.Println("Числа от 1 до 10 в цикле")
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	//Задание 4
	fmt.Println("Введите строку")
	var str string
	fmt.Scanln(&str)
	fmt.Println("Длина строки: ", length(str))

	//Задание 5
	fmt.Println("Введите стороны треугольника")
	var aa, b, c float64
	fmt.Scanln(&aa, &b, &c)
	r := Rectangle{a: aa, b: b, c: c}
	fmt.Println("Площадь ", square(r))

	fmt.Println("aaaa")

}

func isEven(a int) bool {
	if a%2 == 1 {
		return false
	} else {
		return true
	}
}

func check(a int) string {
	if a > 0 {
		return "Positive"
	}
	if a < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func length(str string) int {
	return 0
}

func aver(a, b int) float64 {
	return (float64(a) + float64(b)) / 2.0
}

// Структура треугольник и работа с ней
type Rectangle struct {
	a, b, c float64
}

func newRectangle(a, b, c float64) *Rectangle {
	r := Rectangle{a: a, b: b, c: c}
	return &r
}

func square(r Rectangle) float64 {
	p := r.a + r.b + r.c
	return math.Sqrt(p * (p - r.a) * (p - r.b) * (p - r.c))
}
