package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sebdah/go-jek-battleship/battleground"
	"github.com/sebdah/go-jek-battleship/config"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "input", "", "Input file path")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("No input file specified, please use -input")
	}
}

func main() {
	fileHandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Could not open input file")
	}

	cfg, err := config.FromReader(fileHandle)
	if err != nil {
		log.Fatal("Could not parse input file")
	}

	fmt.Print("Player1\n")
	ground1 := battleground.NewBattleground(cfg.M, &cfg.Player1, &cfg.Player2)
	i := 1
	for _, data := range ground1.Play() {
		fmt.Print(data)

		if i%ground1.M == 0 {
			fmt.Print("\n")
		}

		i++
	}

	fmt.Print("\nPlayer2\n")
	ground2 := battleground.NewBattleground(cfg.M, &cfg.Player2, &cfg.Player1)
	i = 1
	for _, data := range ground2.Play() {
		fmt.Print(data)

		if i%ground2.M == 0 {
			fmt.Print("\n")
		}

		i++
	}

	fmt.Printf("\nP1: %d\n", cfg.Player1.Points())
	fmt.Printf("P2: %d\n", cfg.Player2.Points())
	switch {
	case cfg.Player1.Points() > cfg.Player2.Points():
		fmt.Println("Player 1 wins")
	case cfg.Player1.Points() < cfg.Player2.Points():
		fmt.Println("Player 2 wins")
	default:
		fmt.Println("It's a draw")
	}
}
