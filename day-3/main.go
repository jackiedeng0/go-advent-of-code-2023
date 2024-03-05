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

func symbolIsAdjacent(row int, colStart int, colEnd int, symbolMap map[int]bool)

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

	_, err = recordFile.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Parse numbers in line and check against symbol locations
	i = 0
	// sum := 0
	re := regexp.MustCompile(`\d+`)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		foundLocs := re.FindAllStringIndex(line, -1)
		for _, loc := range foundLocs {
			if (symbolsAdjacent())
			number, err := strconv.Atoi(line[loc[0]:loc[1]])
		}
		i++
	}

	j, _ := strconv.Atoi("32i")
	fmt.Printf("%d\n", j+1)
	//	numStart, numEnd
	//	// Check rows i - 1 to i + 1 from numStart - 1 to numEnd + 1
	//	sum += strconv.Atoi(s)
	//}
}

