// Day 4: Scratchcards
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scoreLine(line string) (uint, error) {
	var score uint = 0
	// Parse out "Card #: "
	numbersStrs := strings.Split(line, ":")
	if len(numbersStrs) < 2 {
		return 0, errors.New("Invalid line")
	}
	// Separate out winning numbers and our numbers by '|' char
	vertSplitStrs := strings.Split(numbersStrs[1], "|")
	if len(vertSplitStrs) < 2 {
		return 0, errors.New("Invalid card")
	}
	// Create map of winning numbers
	winMap := make(map[int]bool)
	winNumStrs := strings.Split(vertSplitStrs[0], " ")
	for _, winNumStr := range winNumStrs {
		if winNumStr == "" {
			// Account for spaces
			continue
		}
		i, err := strconv.Atoi(winNumStr)
		if err != nil {
			return 0, errors.New("Invalid win number")
		}
		winMap[i] = true
	}
	// Iterate through our numbers and check against map
	ourNumStrs := strings.Split(vertSplitStrs[1], " ")
	for _, ourNumStr := range ourNumStrs {
		if ourNumStr == "" {
			// Account for spaces and newlines
			continue
		}
		i, err := strconv.Atoi(ourNumStr)
		if err != nil {
			return 0, errors.New("Invalid card number")
		}
		if !winMap[i] {
			continue
		}
		// Increase score
		if score == 0 {
			score = 1
		} else {
			score *= 2
		}
	}
	return score, nil
}

func main() {
	cardsFile, err := os.Open("cards.txt")
	if err != nil {
		panic(err)
	}
	defer cardsFile.Close()

	totalScore := uint(0)
	reader := bufio.NewReader(cardsFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lineScore, err := scoreLine(strings.TrimRight(line, "\n"))
		if err != nil {
			fmt.Println(err)
		}
		totalScore += lineScore
	}
	fmt.Printf("Total Score: %d\n", totalScore)
}
