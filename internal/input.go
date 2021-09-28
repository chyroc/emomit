package internal

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func Input(title string) (string, error) {
	prompt := promptui.Prompt{
		Label: title,
	}
	command, err := prompt.Run()
	if err != nil {
		return "", err
	}
	command = strings.TrimSpace(command)
	if command == "" {
		return "", fmt.Errorf("Input is empty")
	}
	return command, nil
}
