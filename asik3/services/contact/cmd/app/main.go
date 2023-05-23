package main

import (
	repository ""
	"fmt"
)

func main() {
	_, err := repository.ConnPostgres("localhost", "guest", 5432, "root", "test")
	if err != nil {
		fmt.Errorf("db bad connection: %s", err)
	}
	fmt.Println("db ok")
}
