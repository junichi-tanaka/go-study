package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"log"
	"regexp"
	"flag"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("too few arguments")
		return
	}
	pattern := args[0]
	filename := args[1]

	re, _ := regexp.Compile(pattern)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 1
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		}
		if re.Match(line) {
			fmt.Printf("%d: %s\n", i, string(line))
		}
		i++
	}
}
