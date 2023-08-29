package main

import (
	"fmt"
	"time"
)

func main() {

	// Obter o tempo atual
	currentTime := time.Now()

	// Seta o tempo de respawn do NPC
	alertTime := currentTime.Add(310 * time.Second)

	for {

		// Obter o tempo atual do sistema operacional
		currentTime = time.Now()

		// Verifica se é hora de emitir o alerta
		if currentTime.After(alertTime) {
			fmt.Print(currentTime.Hour(), ":", currentTime.Minute(), ":", currentTime.Second())
			fmt.Println(" Faltam 10 segundos. \a")
			fmt.Println(" Se prepare!! ")
			alertTime = alertTime.Add(300 * time.Second)
		}

		// Aguarda por um curto período antes de verificar novamente
		time.Sleep(1 * time.Second)
	}
}
