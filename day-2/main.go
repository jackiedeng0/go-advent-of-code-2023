// Day 2: Cube Conundrum
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Given a string for one value e.g. "8 blue"
// Separate the int and color value as string
func parseValue(s string) (int, string) {
	subStrs := strings.Split(s, " ")
	count, err := strconv.Atoi(subStrs[0])
	if err != nil {
		log.Fatal(err)
	}
	return count, subStrs[1]
}

// Given a string representing a draw e.g. "8 green, 2 blue"
// Output counts for each color r = 0, g = 8, b = 2
func parseDraw(s string) (r int, g int, b int) {
	r, g, b = 0, 0, 0
	valueStrs := strings.Split(s, ",")
	for i := range len(valueStrs) {
		// Trim out spaces and parseValue
		count, colorStr := parseValue(strings.Trim(valueStrs[i], " "))
		switch colorStr {
		case "red":
			r = count
		case "green":
			g = count
		case "blue":
			b = count
		default:
			fmt.Printf("Invalid colorStr: %s\n", colorStr)
		}
	}
	return
}

// Given a string representing a game e.g. "Game 1: 8 green, 2 blue; 1 red...\n"
// Output the gameID, and the max count for each color in that game
func parseGame(s string) (gameID int, rMax int, gMax int, bMax int) {
	gameID, rMax, gMax, bMax = 0, 0, 0, 0
	subStrs := strings.Split(strings.Trim(s, " \n"), ":")
	if len(subStrs) < 2 {
		fmt.Printf("Bad Game `%s`\n", s)
		return
	}
	gameSubStrs := strings.Split(subStrs[0], " ")
	gameID, err := strconv.Atoi(gameSubStrs[1])
	if err != nil {
		fmt.Printf("Bad GameID `%s`\n Err: %s", s, err)
		return
	}
	// We could check the 'Game' keyword but we'll skip this for succinctness

	// Parse draws for the game
	drawSubStrs := strings.Split(subStrs[1], ";")
	for i := range len(drawSubStrs) {
		r, g, b := parseDraw(drawSubStrs[i])
		if r > rMax {
			rMax = r
		}
		if g > gMax {
			gMax = g
		}
		if b > bMax {
			bMax = b
		}
	}
	return
}

// Takes r g b counts as arguments, reads game.record for the game record and
// prints the sum of possible Game IDs with the counts r, g, b.
func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: <exec> r g b")
	}
	// rCount, _ := strconv.Atoi(os.Args[1])
	// gCount, _ := strconv.Atoi(os.Args[2])
	// bCount, _ := strconv.Atoi(os.Args[3])

	recordFile, err := os.Open("game.record")
	if err != nil {
		log.Fatal(err)
	}
	defer recordFile.Close()

	reader := bufio.NewReader(recordFile)
	// sum := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		gameId, rMax, gMax, bMax := parseGame(line)
		fmt.Println(gameId, rMax, gMax, bMax)
	}
}
