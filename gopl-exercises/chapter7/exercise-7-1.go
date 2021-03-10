package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var fileName = flag.String("name", "", "file to be counted")

func main() {
	flag.Parse()
	fmt.Printf("args fileName is [%s]\n", *fileName)
	if *fileName == "" {
		fmt.Println("fileName cannot be empty! exit...")
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadFile(*fileName)
	if err != nil {
		fmt.Printf("Reading file [%s] get error [%v]\n", *fileName, err)
		os.Exit(2)
	}
	reader := strings.NewReader(string(data))

	scanner := bufio.NewScanner(reader)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		//fmt.Println(scanner.Text())
	}
	//fmt.Println("lineCount = ", lineCount)

	reader = strings.NewReader(string(data))

	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
		//fmt.Println(scanner.Text())
	}
	//fmt.Println("wordCount = ", lineCount)

	// strings.Split(data, "")
	fmt.Printf("LineC = %d, WordC = %d\n", lineCount, wordCount)
}
