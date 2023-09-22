package main

import (
	"fmt"
	"log"
	"monte-carlo-simulation/paytable"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("application has started")

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	paytable.ReadFromFile(os.Getenv("PAYTABLE_FILE_PATH"))

	fmt.Println(paytable.Paytable)
}
