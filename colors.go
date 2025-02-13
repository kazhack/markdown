package main

import "fmt"

// ANSI escape-коды для цветов
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// Функции для цветного текста
func colorText(text, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
