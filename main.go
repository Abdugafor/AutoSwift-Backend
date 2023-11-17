package main

import (
	"AutoSwift/db"
	"log"
)

func main() {
	log.Println("Service launched!")
	db.StartDatabase()
	defer db.CloseDB()
}
