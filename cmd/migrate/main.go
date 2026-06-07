package main

import (
	"flag"
	"log"

	"github.com/v2code/autolog/internal/database"
)

func main() {
	command := flag.String("command", "up", "migration command: up, down, status")
	flag.Parse()

	db := database.OpenDatabase()
	defer db.Close()

	if err := database.RunMigrationCommand(db, *command); err != nil {
		log.Fatal(err)
	}
}
