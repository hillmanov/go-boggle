package main

import (
	"fmt"
	"github.com/hillmanov/go-boggle/board"
	"github.com/hillmanov/go-boggle/dictionary"
	"github.com/hillmanov/go-boggle/word"
)

var d = dictionary.New()
var b = board.New([]string{"b", "e", "r", "t", "g", "h", "n", "m", "a", "e", "w", "p", "l", "e", "t", "f"})
var minWordLength int = 3
var foundWords = make(map[string]bool)
var numberOfChan = 0
var doneChan = make(chan int)

func main() {
	for row := 0; row < b.Size(); row++ {
		for col := 0; col < b.Size(); col++ {
			numberOfChan += 1
			go startFinding(word.New(b.Letter(row, col), row, col), row, col, doneChan)
		}
	}

	for i := 0; i < numberOfChan; i++ {
		<-doneChan
	}

	fmt.Println("Home many words NOW", len(foundWords))
}

func startFinding(w *word.Word, row, col int, doneChan chan <- int) {
	findWords(w, row, col)
	doneChan <- 1
}


func findWords(w *word.Word, row, col int) {
	if isValidWord(w.Word()) {
		foundWords[w.Word()] = true
	}

	validCoordsForWord := getValidCoords(w, row, col)
	for _, coords := range validCoordsForWord {
		if d.ContainsPrefix(w.Word() + b.Letter(coords[0], coords[1])) {
			newWord := word.NewFromWord(w)
			newWord.AddLetter(b.Letter(coords[0], coords[1]), coords[0], coords[1])
			findWords(newWord, coords[0], coords[1])
		}
	}
}

func isValidWord(w string) bool {
	return len(w) >= minWordLength && d.ContainsWord(w)
}

func getValidCoords(w *word.Word, row, col int) [][]int {
	validCoords := make([][]int, 0)

	for loopRow := row - 1; loopRow < row+2; loopRow++ {
		for loopCol := col - 1; loopCol < col+2; loopCol++ {
			if loopRow >= 0 && loopRow < b.Size() && loopCol >= 0 && loopCol < b.Size() {
				if !w.ContainsCoords(loopRow, loopCol) {
					validCoords = append(validCoords, []int{loopRow, loopCol})
				}
			}
		}
	}

	return validCoords
}
