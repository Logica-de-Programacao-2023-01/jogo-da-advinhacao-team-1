package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	rand.Seed(time.Now().UnixNano())
	Resposta := rand.Intn(100)
	for {

		var x int
		fmt.Println("Digite um Numero")
		fmt.Scanln(&x)

		if x == Resposta {
			fmt.Println("certo")
		} else if x < Resposta {
			fmt.Println("o numero e maior")
		} else {
			fmt.Println("O numero e menor")

		}

	}

	jogarnovamente()
}

func jogarnovamente() {
	fmt.Print("Deseja jogar novamente? (sim/não): ")
	var carol string
	fmt.Scanln(&carol)
	carol = strings.ToLower(strings.TrimSpace(carol))
	if carol == "sim" {
		jogarnovamente()
	} else {
		fmt.Println("tks por jogar")
	}
}
