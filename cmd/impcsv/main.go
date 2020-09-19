package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	datas := Prep()
	for i, slice := range datas {
		fmt.Println(i, slice)
	}
	csv := Format(datas)
	for i, v := range csv {
		fmt.Println(i, v)
	}
	recs := cnvToCsv(csv)
	fmt.Println(recs)
}

func Prep() []string {
	s := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}
	return lines
}

func Format(lines []string) []string {
	csv := make([]string, 0)
	beginning := []int{0}
	for i, line := range lines {
		if strings.Contains(line, "class:") {
			beginning = append(beginning, i)
		}
	}
	class := ""
	for i := 0; i < len(lines); i++ {
		for _, v := range beginning {
			if i == v {
				splt := strings.Split(lines[i], " ")
				class = splt[1]
				i++
				break
			}
		}
		newline := class + "," + lines[i]
		csv = append(csv, newline)
	}
	return csv
}

func cnvToCsv(c []string) [][]string {
	recs := make([][]string, 0)
	for _, v := range c {
		line := strings.Split(v, ",")
		recs = append(recs, line)
	}
	return recs
}
