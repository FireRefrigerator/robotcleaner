package main

import (
	"fmt"
	"robotcleaner/app"
	"robotcleaner/domain"
)

func main() {
	fmt.Println("hello cleaner")
	cleaner := domain.NewCleaner(0, 0)
	app.ExecInstraction("./input.json", &cleaner)
}
