package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectAndConvert(input string) (string, error) {
    isMorse := true
    for _, r := range input {
        if !strings.ContainsRune(".- /", r) {
            isMorse = false
            break
        }
    }

    if isMorse {
        result := morse.ToText(input)
        return result, nil
    } else {
        result := morse.ToMorse(input)
        return result, nil
    }
}


