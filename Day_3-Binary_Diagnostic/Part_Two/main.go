package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Report []string

func (r *Report) getFromFile(filePath string) {
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
		*r = append(*r, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type ratingType int

const (
	generator ratingType = iota
	scrubber
)

func (r Report) findRating(t ratingType) int64 {
	for c := 0; c < len(r[0]); c++ {
		var count [2]int
		for _, d := range r {
			r := d[c]
			i := int(r - '0')
			count[i]++
		}
		var priority [2]byte
		switch t {
		case generator:
			priority = [2]byte{'0', '1'}
		case scrubber:
			priority = [2]byte{'1', '0'}
		}
		var (
			more byte
			tmp  Report
		)
		if count[0] > count[1] {
			more = priority[0]
		} else {
			more = priority[1]
		}
		for _, s := range r {
			if s[c] == more {
				tmp = append(tmp, s)
			}
		}
		r = tmp
		if len(r) == 1 {
			break
		}
	}
	result, err := strconv.ParseInt(r[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	var report Report
	report.getFromFile("../input.txt")

	fmt.Println(report.findRating(generator) * report.findRating(scrubber))
}
