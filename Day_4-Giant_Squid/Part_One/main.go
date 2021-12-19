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

func (m Matched) checkForWin() bool {
	var (
		vertical   [N]int
		horizontal [N]int
		diagonal   [2]int
	)

	for i, row := range m {
		for j, b := range row {
			if b {
				vertical[j]++
				horizontal[i]++
				iA, jA := int(math.Abs(float64(i))), int(math.Abs(float64(j)))
				if iA == jA {
					if i == j {
						diagonal[0]++
					} else {
						diagonal[1]++
					}
				}
			}
		}
	}
	for _, v := range vertical {
		if v >= N {
			return true
		}
	}
	for _, h := range horizontal {
		if h >= N {
			return true
		}
	}
	for _, d := range diagonal {
		if d >= N {
			return true
		}
	}
	return false
}

type Bingo struct {
	numbers []int
	boards  []Board
	matched []Matched
}

func (b *Bingo) appendBoard(board *Board) {
	b.boards = append(b.boards, *board)
	b.matched = append(b.matched, Matched{})
}

func (b *Bingo) mark(n int) int {
	for k, board := range b.boards {
		for i, row := range board {
			for j, r := range row {
				if r == n {
					b.matched[k][i][j] = true
				}
				if won := b.matched[k].checkForWin(); won {
					return k
				}
			}
		}
	}
	return -1
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

func (b *Bingo) Play() {
	for _, n := range b.numbers {
		if won := b.mark(n); won >= 0 {
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
	// bingo.getFromFile("test.txt")

	bingo.Play()
}
