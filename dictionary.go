package main

// ResponseCategory структура для категории ответов
type ResponseCategory struct {
	Triggers  []string `json:"triggers"`
	Responses []string `json:"responses"`
}

// Dictionary структура для всего словаря
type Dictionary map[string]ResponseCategory
