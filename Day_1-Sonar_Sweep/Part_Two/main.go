package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getData(filePath string) []int {
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

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func sum(nums ...int) int {
	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

func main() {
	data := getData("../input.txt")

	var result int
	for i := 1; i < len(data); i++ {
		a, b := sum(data[i-1:i+2]...), sum(data[i:i+3]...)
		if b > a {
			result++
		}
	}

	fmt.Println("Answer:", result)
}
