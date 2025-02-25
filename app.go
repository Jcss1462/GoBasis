package main

import (
	"fmt"

	"example.com/m/primerPaquete"
	"example.com/m/user"
)

func main() {
	primerPaquete.MensajeEntrada()
	fmt.Println(primerPaquete.Nombre)
	fmt.Println("Soy el main")

	var newUser *user.User
	var err error
	newUser, err = user.NewUser("Juan Camilo Salazar Serna", 26)

	if err != nil {
		fmt.Println(err)
	} else {
		newUser.PrintUserInfo()
	}

}
