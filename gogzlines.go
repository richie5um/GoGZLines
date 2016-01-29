package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"log"
	"os"
)

func main() {
	fileFlag := flag.String("File", "", "Source File")
	flag.Parse()

	r, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	rz, _ := gzip.NewReader(r)
	defer rz.Close()

	scanner := bufio.NewScanner(rz)
	count := 0

	for scanner.Scan() {
		count++
	}

	log.Printf("Count: %v lines", count)

	os.Exit(0)
}
