package main

import "fmt"

func main() {
	const minutesPerHour = 60
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	hours := totalMinutes / minutesPerHour
	minutes := totalMinutes % minutesPerHour

	fmt.Printf("Получено минут: %d. Добавьте вычисления.\n", totalMinutes)
	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", minutes)
}
