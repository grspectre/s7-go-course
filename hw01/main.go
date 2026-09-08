package main

import (
	"fmt"
	"os"
	"runtime"
)

func getUser() string {
	user := os.Getenv("USERNAME")
	if user == "" {
		user = os.Getenv("USER")
	}
	if user == "" {
		return "Не удалось определить пользователя"
	}
	return user
}

func main() {
	fmt.Println("========================================")
	fmt.Println("             СИСТЕМНЫЙ ОТЧЁТ            ")
	fmt.Println("========================================")

	// 1. Имя текущего пользователя ОС
	fmt.Printf("Пользователь ОС : %s\n", getUser())
	fmt.Println("----------------------------------------")

	// 2. Текущая версия Go
	fmt.Printf("Версия Go       : %s\n", runtime.Version())
	fmt.Println("----------------------------------------")

	// 3. CLI-аргументы
	args := os.Args[1:]
	fmt.Println("Аргументы командной строки:")
	if len(args) == 0 {
		fmt.Println("  (аргументы отсутствуют)")
	} else {
		for i, arg := range args {
			fmt.Printf("  [%d]: %s\n", i+1, arg)
		}
	}
	fmt.Println("========================================")
}
