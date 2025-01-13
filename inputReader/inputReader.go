package inputReader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(promptText string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(promptText)

		enteredText, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read input: %w", err)
		}

		enteredText = strings.TrimSpace(enteredText)

		if enteredText == "" {
			fmt.Println("Input cannot be empty. Please try again.")
			continue
		}

		return enteredText, nil
	}
}
