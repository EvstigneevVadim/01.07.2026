package main

import (
	"fmt"
)

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
--- Меню ---
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
Выберите команду: `

func main() {
	const size = 512
	empls := [size]*Employee{}

	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			empl := new(Employee)
			fmt.Print("Имя: ")
			fmt.Scan(&empl.Name)
			fmt.Print("Возраст: ")
			fmt.Scan(&empl.Age)
			fmt.Print("Позиция: ")
			fmt.Scan(&empl.Position)
			fmt.Print("Зарплата: ")
			fmt.Scan(&empl.Salary)

			added := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					added = true
					fmt.Println("✅ Сотрудник успешно добавлен!")
					break
				}
			}
			if !added {
				fmt.Println("❌ База сотрудников переполнена (достигнут лимит 512)!")
			}

		case 2:
			fmt.Print("Введите имя сотрудника для удаления: ")
			var nameToDelete string
			fmt.Scan(&nameToDelete)

			deleted := false
			for i := 0; i < size; i++ {
				if empls[i] != nil && empls[i].Name == nameToDelete {
					empls[i] = nil
					deleted = true
					fmt.Println("✅ Сотрудник", nameToDelete, "успешно удален!")
					break
				}
			}
			if !deleted {
				fmt.Println("❌ Сотрудник с таким именем не найден.")
			}

		case 3:
			fmt.Println("\n--- Список сотрудников ---")
			count := 0
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Printf("Ячейка %d | Имя: %s, Возраст: %d, Должность: %s, Зарплата: %d\n",
						i, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
					count++
				}
			}
			if count == 0 {
				fmt.Println("Список пуст. Сотрудников пока нет.")
			}
			fmt.Println("--------------------------")

		case 4:
			fmt.Println("Выход из программы...")
			return

		default:
			fmt.Println("❌ Неизвестная команда. Введите число от 1 до 4.")
		}
	}
}
