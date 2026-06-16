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
	if len(os.Args) < 2 || len(os.Args) > 3 {
		log.Fatal("Please provide filename or -h")
	}
	arg := os.Args[1]
	_, err := os.Stat(arg)
	if arg == "-h" || arg == "--help" || arg == "-help" || err != nil {
		fmt.Print("Write a YAML file containing:\nrapid_api_url: <URL>\nrapid_api_headers: \nContent-Type: Application\n<KEY>: <VALUE>\nrapid_api_params:\n- search\n- ?query=example%20query&\n- another=true\n\nRun the program with the path to the file `./sparechange conf.yml` it will use the information to fetch data with given parameters and output via text (1) or markdown (2)")
		return
	}
	if out, err := providers.JSearch(arg); err != nil {
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
