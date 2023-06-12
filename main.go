package main

import (
	"bufio"
	"fileFinderGolang/searcher"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter the text")

		text, _ := reader.ReadString('\n')
		words := strings.Split(text, " ")
		fmt.Println("you entered {}", text)
		s := searcher.Searcher{FileName: words[1], Directory: words[0]}
		err := s.FileSearch()
		if err != nil {
			return
		}
	}
}
