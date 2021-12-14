package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontal, depth int
}

func (p *Position) evalCommand(c Command) {
	switch c.direction {
	case "forward":
		p.horizontal += c.value
	case "up":
		p.depth -= c.value
	case "down":
		p.depth += c.value
	}
}

func (p Position) answer() int {
	return p.depth * p.horizontal
}

type Command struct {
	direction string
	value     int
}

type Commands []Command

func (c *Commands) getFromFile(filePath string) {
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		number, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		*c = append(*c,
			Command{
				direction: line[0],
				value:     number,
			},
		)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	data := Commands{}
	data.getFromFile("../input.txt")

	position := Position{}

	for _, d := range data {
		position.evalCommand(d)
	}

	fmt.Println(position.answer())
}
