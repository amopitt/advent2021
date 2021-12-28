package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 21")
	fmt.Println(input)
	fmt.Println("------------------")

	var player1 Player = Player{
		currentSpace: 10,
		totalScore:   0,
	}
	var player2 Player = Player{
		currentSpace: 16,
		totalScore:   0,
	}
	fmt.Println("Part 1:", part1(player1, player2))
}

func part1(player1 Player, player2 Player) int {

	rolls := 3
	die := 1
	turn := 1

	totalRolls := 0

	for {
		nextRoll, total := roll(die, rolls)
		die = nextRoll
		totalRolls += rolls

		if turn%2 == 0 {
			player2.move(total)
			fmt.Println("Player 2: currentSpace", player2.currentSpace, " total=", player2.totalScore)

		} else {
			player1.move(total)
			fmt.Println("Player 1: currentSpace", player1.currentSpace, " total=", player1.totalScore)

		}
		turn++

		if player1.totalScore >= 1000 {
			return player2.totalScore * totalRolls
		}

		if player2.totalScore >= 1000 {

			return player2.totalScore * totalRolls
		}

		if die > 100 {
			break
		}
	}
	return 0

}
func roll(die int, count int) (int, int) {
	total := 0
	nextRoll := die

	for i := 1; i <= count; i++ {
		fmt.Println("rolling", "d", i, nextRoll)
		total += nextRoll
		nextRoll++
		if nextRoll == 101 {
			nextRoll = 1
		}
	}
	return nextRoll, total
}

func (p *Player) move(diceRoll int) {
	p.currentSpace += diceRoll

	// wrap around space   from 1 to 10
	p.currentSpace = p.currentSpace % 10
	if p.currentSpace == 0 {
		p.currentSpace = 10
	}
	p.totalScore += p.currentSpace
}

type Player struct {
	currentSpace int
	totalScore   int
}
