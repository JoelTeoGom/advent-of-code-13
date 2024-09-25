package main

import (
	"bufio"
	"fmt"
	"os"
)

type Puzzle struct {
	data       [][]string
	Row, Col   int
	Reflection int
}

// Function to check if one row or column is a mirror of another
func isMirror(row1, row2 []string) bool {
	fmt.Println("Row 1: ", row1)
	fmt.Println("row 2: ", row2)
	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}

// Function to find mirrors in rows (horizontal)
func horizontal(puzzle *Puzzle, sumarize *int, nPuzzle int) {
	fmt.Println("Horizontal [", nPuzzle, "]")

	for i := 0; i < puzzle.Row-1; i++ { // -1 to avoid index out of range
		//fmt.Println(puzzle.data[i])
		if isMirror(puzzle.data[i], puzzle.data[i+1]) {
			puzzle.Reflection = i + 1
			fmt.Printf("Found a mirror in rows %d and %d\n", i, i+1)
		}
	}

	*sumarize = *sumarize + (puzzle.Reflection * 100)
	//fmt.Println("Suma horizontal: ", *sumarize)
}

// Function to find mirrors in columns (vertical)
func vertical(puzzle *Puzzle, sumarize *int, nPuzzle int) {
	column := make([][]string, puzzle.Col)
	fmt.Println("Vertical[", nPuzzle, "]")

	// Transpose the puzzle data into columns
	for i := 0; i < puzzle.Row; i++ {
		for j := 0; j < puzzle.Col; j++ {
			column[j] = append(column[j], puzzle.data[i][j])
		}
	}

	puzzle.Reflection = 0
	// Compare columns to find mirrors
	for i := 0; i < puzzle.Col-1; i++ { // -1 to avoid index out of range
		//fmt.Println(column[i])
		if isMirror(column[i], column[i+1]) {
			puzzle.Reflection = i + 1
			fmt.Printf("Found a mirror in columns %d and %d\n", i, i+1)
		}
	}
	*sumarize += puzzle.Reflection

	//fmt.Println("Suma vertical: ", *sumarize)
}

func readfile(sumarize *int, puzzle *Puzzle, filename string) {
	var nPuzzle = 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" { // Check for empty lines
			//fmt.Println("Puzzle data before reset:", puzzle.data)
			vertical(puzzle, sumarize, nPuzzle)
			horizontal(puzzle, sumarize, nPuzzle)

			// Reset the puzzle for the next segment
			puzzle.data = nil
			puzzle.Row = 0
			puzzle.Col = 0
			puzzle.Reflection = 0
			nPuzzle++
			continue // Skip to the next iteration
		}

		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}

		puzzle.Col = len(row)                  // Set column count based on row length
		puzzle.Row += 1                        // Increment row count
		puzzle.data = append(puzzle.data, row) // Append the row to puzzle data
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//for the last
	vertical(puzzle, sumarize, nPuzzle)
	horizontal(puzzle, sumarize, nPuzzle)
	nPuzzle++
	fmt.Println("npuzzle: ", nPuzzle)

}

func main() {
	// Example usage
	puzzle := Puzzle{}
	var sumarize int
	readfile(&sumarize, &puzzle, "puzzle.txt") // Make sure to have this file in the directory
	fmt.Println(sumarize)
}
