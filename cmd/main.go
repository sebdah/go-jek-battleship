package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sebdah/go-jek-battleship/battleground"
	"github.com/sebdah/go-jek-battleship/config"
)

var (
	inputFile  string
	outputFile string
)

func init() {
	flag.StringVar(&inputFile, "input", "", "Input file path")
	flag.StringVar(&outputFile, "output", "output.txt", "Output file path")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("No input file specified, please use -input")
	}

	if outputFile == "" {
		log.Fatal("No input file specified, please use -output")
	}
}

func main() {
	inputFileHandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Could not open input file")
	}
	defer inputFileHandle.Close()

	cfg, err := config.FromReader(inputFileHandle)
	if err != nil {
		log.Fatal("Could not parse input file")
	}

	err = cfg.Validate()
	if err != nil {
		log.Fatalf("Invalid configuration: %s", err.Error())
	}

	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Could not open output file")
	}
	defer outputFileHandle.Close()

	ground := battleground.NewBattleground(cfg.M, cfg.Player1, cfg.Player2)
	playBoard("Player1", ground, outputFileHandle)
	outputFileHandle.WriteString("\n")
	ground = battleground.NewBattleground(cfg.M, cfg.Player2, cfg.Player1)
	playBoard("Player2", ground, outputFileHandle)

	outputFileHandle.WriteString(fmt.Sprintf("\nP1: %d\n", cfg.Player1.Points()))
	outputFileHandle.WriteString(fmt.Sprintf("P2: %d\n", cfg.Player2.Points()))
	switch {
	case cfg.Player1.Points() > cfg.Player2.Points():
		outputFileHandle.WriteString("Player 1 wins\n")
	case cfg.Player1.Points() < cfg.Player2.Points():
		outputFileHandle.WriteString("Player 2 wins\n")
	default:
		outputFileHandle.WriteString("It's a draw\n")
	}
}

func playBoard(name string, battleground battleground.Battlegrounder, out *os.File) {
	out.WriteString(fmt.Sprintf("%s\n", name))
	i := 1
	for _, data := range battleground.Play() {
		out.WriteString(data)

		if i%battleground.M() == 0 {
			out.WriteString("\n")
		}

		i++
	}
}
