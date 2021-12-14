package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getData(filePath string) []string {
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

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	data := getData("../input.txt")

	var gamma string
	for c := 0; c < len(data[0]); c++ {
		var count [2]int
		for _, d := range data {
			r := d[c]
			i := int(r - '0')
			count[i]++
		}
		if count[0] > count[1] {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}

	var epsilon string
	for _, r := range gamma {
		switch r {
		case '0':
			epsilon += "1"
		case '1':
			epsilon += "0"
		}

	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e * g)
}
