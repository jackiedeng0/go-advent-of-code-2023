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
func parseValue(s string) (c int, colorStr string) {
	c, colorStr = 0, ""
	_, err := fmt.Sscanf(s, "%d %s", &c, &colorStr)
	if err != nil {
		log.Fatal(err)
	}
	return
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
	subStrs := strings.Split(s, ": ")
	if len(subStrs) < 2 {
		fmt.Printf("Bad Game format `%s`\n", s)
	}
	n, err := fmt.Sscanf(subStrs[0], "Game %d", &gameID)
	if err != nil {
		fmt.Printf("Bad Game format `%s`\n", s)
		fmt.Println(n, err)
		return
	}

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
	rCount, _ := strconv.Atoi(os.Args[1])
	gCount, _ := strconv.Atoi(os.Args[2])
	bCount, _ := strconv.Atoi(os.Args[3])

	recordFile, err := os.Open("game.record")
	if err != nil {
		log.Fatal(err)
	}
	defer recordFile.Close()

	scanner := bufio.NewScanner(recordFile)
	sum := 0
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)

		gameId, rMax, gMax, bMax := parseGame(line)
		fmt.Println(gameId, rMax, gMax, bMax)

		if (rMax > rCount) || (gMax > gCount) || (bMax > bCount) {
			fmt.Println("Game not possible!")
			continue
		}
		sum += gameId
	}
	fmt.Println("Sum of Game IDs: ", sum)
}
