package main

import "fmt"

func main() {

	const course = "Бэкенд-разработка на языке Go"
	var name string = "Мартин"
	var group string = "Учебная-01"
	completed := 1
	var ready bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Println("Курс: ", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", ready)
}
