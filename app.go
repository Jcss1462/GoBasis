package main

import (
	"fmt"

	"example.com/m/primerPaquete"
)

func main() {
	primerPaquete.MensajeEntrada()
	fmt.Println(primerPaquete.Nombre)
	fmt.Println("Soy el main")
}
