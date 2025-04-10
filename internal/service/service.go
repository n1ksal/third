package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

type Converter struct{}

// Convert принимает входную строку и определяет, необходимо ли преобразовать ее в Морзе или обратно.
func (c Converter) Convert(input string) (string, error) {
	input = strings.TrimSpace(input)

	// Проверка на наличие символов Морзе в вводе
	if strings.ContainsAny(input, ".-") {
		return c.convertMorseToText(input)
	}
	return c.convertTextToMorse(input)
}

// convertMorseToText преобразует код Морзе в текст
func (c Converter) convertMorseToText(morseCode string) (string, error) {
	text := morse.ToText(morseCode)
	if text == "" {
		return "", errors.New("не удалось декодировать код Морзе")
	}
	return text, nil
}

// convertTextToMorse преобразует текст в код Морзе
func (c Converter) convertTextToMorse(text string) (string, error) {
	var morseCode strings.Builder

	for _, r := range text {
		morseChar := morse.RuneToMorse(r)
		if morseChar == "" {
			return "", errors.New("не удалось закодировать символ: " + string(r))
		}
		morseCode.WriteString(morseChar + " ")
	}

	return strings.TrimSpace(morseCode.String()), nil
}
