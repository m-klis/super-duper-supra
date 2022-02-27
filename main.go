package main

import (
	"noteapp/config"
)

func main() {
	// setup config
	configuration := config.New()
	db := config.NewDatabase(configuration)
	// fmt.Println(db)
}
