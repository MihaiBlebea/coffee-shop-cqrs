package main

import (
	"fmt"

	"github.com/MihaiBlebea/coffee-shop-query/cmd"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Execute()
}
