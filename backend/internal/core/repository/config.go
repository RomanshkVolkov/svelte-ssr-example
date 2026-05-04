package repository

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RomanshkVolkov/svelte-and-go-template/internal/core/lg"
)

func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("No .env file found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.ReplaceAll(strings.TrimSpace(parts[1]), "\"", "")
			if os.Getenv(key) == "" { // No sobreescribir variables del sistema
				lg.Info("Variable " + key + " cargada con valor " + value)
				os.Setenv(key, value)
			}
		}
	}
}

func GetEnv(key string, defaultValue string) string {
	variable := os.Getenv(key)

	if variable == "" {
		fmt.Println("The environment variable " + key + " is not set. Using default value.")
		return defaultValue
	}

	return variable
}
