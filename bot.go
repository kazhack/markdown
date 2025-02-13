package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

// Bot структура для бота
type Bot struct {
	Dict Dictionary
}

// NewBot создает нового бота из JSON файла
func NewBot(jsonPath string) (*Bot, error) {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var dict Dictionary
	if err := json.Unmarshal(data, &dict); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &Bot{Dict: dict}, nil
}

// findCategory ищет подходящую категорию для сообщения
func (b *Bot) findCategory(message string) string {
	message = strings.ToLower(strings.TrimSpace(message))

	for category, data := range b.Dict {
		for _, trigger := range data.Triggers {
			if strings.Contains(message, strings.ToLower(trigger)) {
				return category
			}
		}
	}
	return "unknown"
}

// GetResponse возвращает случайный ответ для сообщения
func (b *Bot) GetResponse(message string) string {
	category := b.findCategory(message)

	var responses []string
	if category == "unknown" {
		responses = b.Dict["unknown"].Responses
	} else {
		responses = b.Dict[category].Responses
	}

	if len(responses) == 0 {
		return "Извините, я не знаю, что ответить."
	}

	return responses[rand.Intn(len(responses))]
}
