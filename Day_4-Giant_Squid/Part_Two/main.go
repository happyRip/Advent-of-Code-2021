package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const N = 5

type Board [N][N]int

type Matched [N][N]bool

type Scores struct {
	vertical   [N]int
	horizontal [N]int
	diagonal   [2]int
}

func (s *Scores) increment(i, j int) bool {
	if s.horizontal[i]++; s.horizontal[i] >= N {
		return true
	}
	if s.vertical[j]++; s.vertical[j] >= N {
		return true
	}
	iA, jA := int(math.Abs(float64(i))), int(math.Abs(float64(j)))
	if iA == jA {
		if i == j {
			s.diagonal[0]++
		} else {
			s.diagonal[1]++
		}
		if s.diagonal[0] >= N || s.diagonal[1] >= N {
			return true
		}
	}
	return false
}

type Winners struct {
	won []bool
	n   int
}

type Bingo struct {
	numbers []int
	boards  []Board
	matched []Matched
	scores  []Scores
	winners Winners
}

func (b *Bingo) mark(k, i, j int) {
	toMatch := &b.matched[k][i][j]
	if !*toMatch {
		if b.scores[k].increment(i, j) {
			if !b.winners.won[k] {
				b.winners.n++
			}
			b.winners.won[k] = true
		}
		*toMatch = true
	}
}

func (b *Bingo) processNumber(n int) int {
	for k, board := range b.boards {
		if !b.winners.won[k] {
			for i, row := range board {
				for j, r := range row {
					if r == n {
						b.mark(k, i, j)
						if b.winners.n >= len(b.boards) {
							return k
						}
					}
				}
			}
		}
	}
	return -1
}

func (b *Bingo) play() {
	for _, n := range b.numbers {
		won := b.processNumber(n)
		if won != -1 {
			unmarked := b.sumUnmarked(won)
			fmt.Println(unmarked * n)
			return
		}
	}
}

func (b Bingo) sumUnmarked(n int) int {
	var unmarked int
	board, matched := b.boards[n], b.matched[n]
	for i, row := range board {
		for j, v := range row {
			if !matched[i][j] {
				unmarked += v
			}
		}
	}
	return unmarked
}

func (b *Bingo) appendBoard(board *Board) {
	b.boards = append(b.boards, *board)
	b.matched = append(b.matched, Matched{})
	b.scores = append(b.scores, Scores{})
	b.winners.won = append(b.winners.won, false)
}

func (b *Bingo) getFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var (
		board Board
		i     int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(b.numbers) == 0 {
			b.numbers = lineToIntSlice(line, ",")
		} else {
			if line == "" && i > 0 {
				b.appendBoard(&board)
				board, i = Board{}, 0
			} else if line != "" && len(line) > 0 {
				l := lineToIntSlice(line, " ")
				for j := range board[i] {
					board[i][j] = l[j]
				}
				i++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(board) > 0 {
		b.appendBoard(&board)
	}
}
func lineToIntSlice(line string, separator string) []int {
	nums := strings.Split(line, separator)
	var numbers []int
	for _, n := range nums {
		if n == " " || n == "" {
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}
	return numbers
}

func main() {
	var bingo Bingo
	bingo.getFromFile("../input.txt")

	bingo.play()
}
