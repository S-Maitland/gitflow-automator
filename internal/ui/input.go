package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// PromptString prompts the user for a string input
func PromptString(prompt string) (string, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// PromptInt prompts the user for yes/no confirmation
func PromptYesNo(prompt string) bool {
	reader:= bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " (y/n): ")

	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	return input == "y" || input == "yes"
}

// PromptSelect shows a list and asks user to select one
func PromptSelect(prompt string, options []string) (int, string, error) {
	fmt.Println(prompt)
	for i, option := range options {
		fmt.Printf(" %d) %s\n", i+1, option)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter number: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, "", err
	}

	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(options) {
		return 0, "", fmt.Errorf("invalid selection")
	}

	return choice - 1, options[choice-1], nil
}