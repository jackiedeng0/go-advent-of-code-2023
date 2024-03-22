// Day 3: Gear Ratios
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const symbols string = "!@#$%^&*()-=_+[]\\{}|:';\",/<>?"

// Returns column-indexed map of bools representing whether a symbol exists
// at the column for a given line
func symbolIndicesInLine(line string) map[int]bool {
	lastSymbolIndex := -1
	lineMap := make(map[int]bool)
	for {
		foundSymbolIndex := strings.IndexAny(line[lastSymbolIndex+1:], symbols)
		if foundSymbolIndex == -1 {
			break
		}
		// foundSymbolIndex is relative to lastSymbolIndex + 1 so we need add
		// that to update lastSymbolIndex accordingly
		lastSymbolIndex += 1 + foundSymbolIndex
		lineMap[lastSymbolIndex] = true
	}
	return lineMap
}

// Returns whether there is any symbol adjacent to the location of a number
func symbolIsAdjacent(row int, lastRow int, startCol int, endCol int,
	symbolMap map[int]map[int]bool) bool {
	// Check current row
	// Note that if startCol is 0 or endCol is the line length, this makes an
	// out of bounds access but go maps are zero initialized so this is
	// actually ok
	if symbolMap[row][startCol-1] || symbolMap[row][endCol] {
		return true
	}
	if row >= 1 {
		// Check row above
		for c := startCol - 1; c <= endCol; c++ {
			if symbolMap[row-1][c] {
				return true
			}
		}
	}
	if row < lastRow {
		// Check row below
		for c := startCol - 1; c <= endCol; c++ {
			if symbolMap[row+1][c] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Parse symbol locations as a row-indexed map of column-indexed maps
	symbolsMap := make(map[int]map[int]bool)

	recordFile, err := os.Open("engine.sch")
	if err != nil {
		panic(err)
	}
	defer recordFile.Close()

	i := 0
	reader := bufio.NewReader(recordFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		symbolsMap[i] = symbolIndicesInLine(line)
		i++
	}
	totalRows := i

	_, err = recordFile.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Parse numbers in line and check against symbol locations
	i = 0
	sum := 0
	re := regexp.MustCompile(`\d+`)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		foundLocs := re.FindAllStringIndex(line, -1)
		for _, loc := range foundLocs {
			if symbolIsAdjacent(i, totalRows, loc[0], loc[1], symbolsMap) {
				number, err := strconv.Atoi(line[loc[0]:loc[1]])
				if err != nil {
					panic(err)
				}
				sum += number
			}
		}
		i++
	}

	fmt.Printf("Sum: %d\n", sum)
}
