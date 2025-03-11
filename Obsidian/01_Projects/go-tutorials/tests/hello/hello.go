package hello

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empte name")
	}
	greeting := fmt.Sprintf("Hello %v How are you ?", name)
	return greeting, nil
}
