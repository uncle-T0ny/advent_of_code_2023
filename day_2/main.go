package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID         int
	Sequences  []Sequence
	IsPossible bool
}

type Sequence struct {
	Loads []Load
}

type Load struct {
	Amount int
	Color  string
}

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumPart1 := 0
	sumPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)

		if game.IsPossible {
			sumPart1 += game.ID
		}

		sumPart2 += getGameFewestNumPow(game)
	}
	fmt.Println("part 1 res:", sumPart1)
	fmt.Println("part 2 res:", sumPart2)
}

func getGameFewestNumPow(game Game) int {
	game.IsPossible = true
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, seq := range game.Sequences {
		for _, load := range seq.Loads {
			if load.Color == "red" && load.Amount > maxRed {
				maxRed = load.Amount
			}
			if load.Color == "green" && load.Amount > maxGreen {
				maxGreen = load.Amount
			}
			if load.Color == "blue" && load.Amount > maxBlue {
				maxBlue = load.Amount
			}
		}
	}

	return maxRed * maxGreen * maxBlue
}

func parseGame(line string) Game {
	parts := strings.Split(line, ":")
	fmt.Println(parts)

	gameRaw := strings.TrimSpace(parts[0])
	gameIDStr := strings.Split(gameRaw, " ")[1]
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sequences := strings.Split(parts[1], ";")

	game := Game{
		ID:         gameID,
		Sequences:  []Sequence{},
		IsPossible: true,
	}

	for _, seq := range sequences {
		seq = strings.TrimSpace(seq)
		loading := strings.Split(seq, ",")

		loads := []Load{}
		for _, load := range loading {
			load = strings.TrimSpace(load)
			amount := strings.Split(load, " ")[0]
			color := strings.Split(load, " ")[1]

			amountNum, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			loads = append(loads, Load{
				Amount: amountNum,
				Color:  color,
			})

			if color == "red" && amountNum > 12 {
				game.IsPossible = false
			}
			if color == "green" && amountNum > 13 {
				game.IsPossible = false
			}
			if color == "blue" && amountNum > 14 {
				game.IsPossible = false
			}
		}

		game.Sequences = append(game.Sequences, Sequence{
			Loads: loads,
		})
	}

	return game
}
