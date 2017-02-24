package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/sebdah/go-jek-battleship/source/battleground"
	"github.com/sebdah/go-jek-battleship/source/coordinates"
)

var inputFile string

// NONE is for unused postitions.
const NONE = "_"

// MISS is for missile misses.
const MISS = "O"

// HIT is for hit battleships.
const HIT = "X"

// ALIVE is for alive battleships.
const ALIVE = "B"

func init() {
	flag.StringVar(&inputFile, "input", "", "Input file path")
	flag.Parse()

	if inputFile == "" {
		logrus.Fatal("No input file specified, please use -input")
	}
}

func main() {
	fileHandle, err := os.Open(inputFile)
	if err != nil {
		logrus.WithError(err).Fatal("Could not open input file")
	}

	config, err := battleground.ConfigFromReader(fileHandle)
	if err != nil {
		logrus.WithError(err).Fatal("Could not parse input file")
	}
	play(config)
}

func play(cfg *battleground.Config) {
	matrix := []string{}

	for y := 0; y < cfg.M; y++ {
		for x := 0; x < cfg.M; x++ {
			coordinate := coordinates.NewCoordinate(x, y)

			isMissile := false
			if cfg.Player1.Moves.Include(coordinate) || cfg.Player2.Moves.Include(coordinate) {
				isMissile = true
			}

			isShip := false
			if cfg.Player1.ShipPositions.Include(coordinate) || cfg.Player2.ShipPositions.Include(coordinate) {
				isShip = true
			}

			switch {
			case !isMissile && isShip:
				matrix = append(matrix, ALIVE)
			case isMissile && isShip:
				matrix = append(matrix, HIT)
			case isMissile && !isShip:
				matrix = append(matrix, MISS)
			default:
				matrix = append(matrix, NONE)
			}
		}
	}

	i := 1
	for _, data := range matrix {
		fmt.Print(data)

		if i%cfg.M == 0 {
			fmt.Print("\n")
		}

		i++
	}
}
