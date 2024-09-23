package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
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
		fmt.Print(i, " ")
	}
	fmt.Println()

	//Задание 4
	fmt.Println("Введите строку")
	var str string
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = text
	//fmt.Scan(&str)
	fmt.Println("Длина строки: ", length(str))

	//Задание 5
	fmt.Println("Введите стороны прямоугольника")
	var aa, b float64
	fmt.Scanln(&aa, &b)
	r := Rectangle{a: aa, b: b}
	fmt.Println("Площадь ", square(r))

	//Задание 6
	fmt.Println("Введите два целых числа")
	var bb int
	fmt.Scanln(&a, &bb)
	fmt.Println("Среднее двух чисел: ", aver(a, bb))

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
	return utf8.RuneCountInString(str)
}

func aver(a, b int) float64 {
	return (float64(a) + float64(b)) / 2.0
}

// Структура треугольник и работа с ней
type Rectangle struct {
	a, b float64
}

func newRectangle(a, b, c float64) *Rectangle {
	r := Rectangle{a: a, b: b}
	return &r
}

func square(r Rectangle) float64 {
	return r.a * r.b
}
