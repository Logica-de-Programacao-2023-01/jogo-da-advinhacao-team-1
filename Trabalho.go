package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var playAttempts []int

func main() {

	fmt.Println("Bem-vindo ao jogo da adivinhação!")
	play()

}

func play() {
	num := rand.Intn(100)
	var attempts int
	for {

		var entry int
		fmt.Println("Digite um número entre 1 e 100:")
		_, err := fmt.Scan(&entry)
		if err != nil {
			fmt.Println("Digite um número valido.")
		} else if entry == num {
			fmt.Println("Parabéns, você acertou!")
			break
		} else if entry < num {
			fmt.Printf("O número é maior que %d\n", entry)
		} else {
			fmt.Printf("O número é menor que %d\n", entry)
		}

		attempts++

	}

	fmt.Printf("Você utilizou %d tentativas\n", attempts)
	playAttempts = append(playAttempts, attempts)

	playAgain()
}

func playAgain() {
	fmt.Println("Deseja jogar novamente? (sim/não): ")

	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println("Desculpe, um erro ocorreu! :/")
		return
	}
	answer = strings.ToLower(strings.TrimSpace(answer))

	if answer == "sim" || answer == "s" || answer == "yes" || answer == "si" || answer == "y" {
		play()
	} else {
		fmt.Println("\nObrigado por jogar :)")
		for i, attempt := range playAttempts {
			fmt.Printf("Você utilizou %d tentativas na %d jogada\n", attempt, i+1)
		}
	}
}
