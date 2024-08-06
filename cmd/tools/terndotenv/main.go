package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	fmt.Println("WEB_SOCKET_PORT:", os.Getenv("WEB_SOCKET_PORT"))
	fmt.Println("DATABASE_NAME:", os.Getenv("DATABASE_NAME"))
	fmt.Println("DATABASE_USER:", os.Getenv("DATABASE_USER"))
	fmt.Println("DATABASE_PASSWD:", os.Getenv("DATABASE_PASSWD"))
	fmt.Println("DATABASE_HOST:", os.Getenv("DATABASE_HOST"))

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command execution failed with error: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return
	}

	fmt.Printf("Command executed successfully: %s\n", string(output))
}
