package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := os.Args[1]

	// Open the input file for reading
	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Open the output file for appending
	outputFile, err := os.OpenFile("extracted_domains.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Define a regular expression to match URLs
	re := regexp.MustCompile(`\bhttps?://\S+\b`)

	// Use a map to keep track of unique URLs
	uniqueURLs := make(map[string]bool)

	// Read the input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all URLs in the line
		urls := re.FindAllString(line, -1)

		// Write unique URLs to the output file
		for _, url := range urls {
			if !uniqueURLs[url] {
				fmt.Fprintln(outputFile, url)
				uniqueURLs[url] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done extracting unique URLs to extracted_domains.txt")
}



