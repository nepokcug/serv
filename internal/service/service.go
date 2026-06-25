package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const russianLetters = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

func AutoConvert(input string) (string, error) {
	if input == "" {
		return "", errors.New("input string is empty")
	}
	//Если есть хотя бы один символ
	if strings.ContainsAny(input, russianLetters) {
		result := morse.ToMorse(input)
		if result == "" {
			return "", errors.New("failed to convert text to morse code")
		}
		return result, nil
	}
	//Иначе Морзе
	result := morse.ToText(input)
	if result == "" {
		return "", errors.New("failed to convert morse code to text")
	}
	return result, nil
}
