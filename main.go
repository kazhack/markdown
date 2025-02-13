package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем бота
	bot, err := NewBot("data/dictionary.json")
	if err != nil {
		fmt.Printf("Error initializing bot: %v\n", err)
		return
	}

	// Выводим приветствие с цветом для бота
	fmt.Println(colorText("MarkBot запущен! Я говорю только на русском! (Для выхода введите 'выход')", Cyan))
	fmt.Println("----------------------------------------")

	// Инициализируем сканер для чтения ввода с пробелами
	scanner := bufio.NewScanner(os.Stdin)

	// Основной цикл чата
	for {
		// Выводим имя пользователя с цветом
		fmt.Print(colorText("[Пользователь]: ", Green))

		// Считываем ввод пользователя
		if scanner.Scan() {
			input := scanner.Text()

			if strings.ToLower(input) == "выход" {
				// Выводим прощание с цветом
				fmt.Println(colorText("Пошёл на*уй", Yellow))
				break
			}

			// Получаем ответ от бота и выводим его с цветом
			response := bot.GetResponse(input)
			fmt.Printf(colorText("[MarkBot]: %s\n", Blue), response)
		} else {
			fmt.Println("Ошибка при чтении ввода")
			continue
		}
	}
}
