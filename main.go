package main

import (
	"fmt"

	"github.com/csepulvedaz/meli-challenge/src/core"
)

type Satellite struct {
	Name     string
	Distance float32
	Message  []string
}

func main() {
	kenobi := Satellite{"kenobi", 880, []string{"este", "", "", "mensaje", ""}}
	skywalker := Satellite{"skywalker", 790, []string{"", "es", "", "", "secreto"}}
	sato := Satellite{"sato", 883, []string{"este", "", "un", "", ""}}

	message := GetMessage(kenobi.Message, skywalker.Message, sato.Message)
	fmt.Printf("Mesaje: %s\n", message)

	x, y := GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)
	fmt.Printf("Posición: %f, %f\n", x, y)

	core.StartServer()
}
