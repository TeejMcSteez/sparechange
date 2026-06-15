package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sparechange/parser"
	"sparechange/providers"
	"strings"
	"time"
)

func main() {
	conf_file := os.Args[1]
	if len(conf_file) == 0 {
		log.Fatal("Failed to provide configuration file for variables")
	}
	if out, err := providers.JSearch(conf_file); err != nil {
		log.Fatalf("Failed to search JSearch Provider: %v", err)
	} else {
		fmt.Print("Choose output format: \n1: Text File\n2: Markdown\n")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()

			command := strings.TrimSpace(input)
			timestamp_filename := time.Now().String() + "_JSearch_Output"
			switch command {
			case "1":
				if err := parser.WriteText(timestamp_filename+".txt", out); err != nil {
					log.Fatalf("Failed to write output text file: %v", err)
				}
			case "2":
				// Implement markdown ProviderOutput Parser
				if err := parser.WriteMarkdown(timestamp_filename+".md", out); err != nil {
					log.Fatalf("Failed to write output to markdown file: %v", err)
				}
			default:
				log.Fatal("Invalid choice")
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
	}
}
