package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println(
		"Um número aleatório será sorteado. Tente acertar. O número é um inteiro entr 1 a 100",
	)

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)

	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O seu chute deve ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número secreto é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número secreto é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns você acertou o número que era: %d\n"+
					"Você acertou em %d tentativas.\n"+
					"Essas foram suas tentativas: %v\n",
				x, i+1, chutes,
			)
			return
		}

		chutes[i] = chuteInt

	}

	fmt.Printf(
		"Infelizmente você não acertou o número, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram suas tentativas: %v\n",
		x, chutes,
	)
}
