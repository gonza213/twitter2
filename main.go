package main

import (
	"log"

	"github.com/gonza213/twitter2/bd"
	"github.com/gonza213/twitter2/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión en la BD...")
		return
	}

	handlers.Manejadores()
}
