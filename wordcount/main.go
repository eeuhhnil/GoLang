package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filenames := os.Args[1:]

	counts := map[string]int{}

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			counts[scanner.Text()]++
		}
	}
	for line, count := range counts {
		if count >= 2 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
