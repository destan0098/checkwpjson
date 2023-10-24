package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/destan0098/checkwpjson/cmd/Part1"
	"github.com/destan0098/checkwpjson/cmd/Part2"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

var path string

var outputs, inputs, domains string = "", "", ""

var pipel int
var results []DomainResult

type DomainResult struct {
	Domain   string
	UserName string
}

func writeResults(results []DomainResult, outputfile string) {

	file, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// Write the results to a CSV file.
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{"Domain", "UserName"})
	if err != nil {
		return
	}
	for _, result := range results {
		err = writer.Write([]string{result.Domain, result.UserName})
		if err != nil {
			return
		}
	}
}

// readDomains reads domain names from a text file and returns them as a string slice.
func readDomains(filename string) []string {
	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var domainss []string
	var InputText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		InputText = scanner.Text()

		if !strings.HasPrefix(path, "https://") {
			if !strings.HasPrefix(path, "http://") {
				path = "https://" + path

			}

		}
		// Add domain names to the list.
		domainss = append(domainss, InputText)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return domainss
}

func withlist(inputfile string) []DomainResult {
	domainadd := readDomains(inputfile)
	var results []DomainResult
	var result DomainResult
	for _, domain := range domainadd {

		// Check if port 80 is open for the domain.
		users, dAdress := Part1.Part1(domain)
		if users == nil || dAdress == "" {
			users, dAdress = Part2.Part2(domains)
		}
		for _, user := range users {
			result = DomainResult{
				Domain:   dAdress,
				UserName: user,
			}
			results = append(results, result)
		}

	}
	return results
}
func withname(domain string) []DomainResult {

	users, dAdress := Part1.Part1(domain)

	if users == nil || dAdress == "" {
		users, dAdress = Part2.Part2(domains)
	}
	result := DomainResult{
		Domain:   dAdress,
		UserName: users[0],
	}
	results = append(results, result)

	// If port 80 or 443 is open, print a message and store the result.

	fmt.Printf(color.Colorize(color.Green, "[+] Domain %s is Opened\n"), result.Domain)
	fmt.Printf(color.Colorize(color.Green, "[*] In single-domain mode, the output file is not saved.\n"))
	return results

}
func withpip() []DomainResult {
	scanner := bufio.NewScanner(os.Stdin)
	var results []DomainResult
	for scanner.Scan() {

		domain := scanner.Text()

		// Check if port 80 is open for the domain.
		var results []DomainResult
		var result DomainResult

		// Check if port 80 is open for the domain.
		users, dAdress := Part1.Part1(domain)

		if users == nil || dAdress == "" {
			users, dAdress = Part2.Part2(domains)
		}
		for _, user := range users {
			result = DomainResult{
				Domain:   dAdress,
				UserName: user,
			}
			results = append(results, result)
		}

		// If port 80 or 443 is open, print a message and store the result.

		fmt.Printf(color.Colorize(color.Green, "[+] Domain %s is Opened\n"), result.Domain)
		results = append(results, result)

	}

	return results
}
func main() {
	//Receive input from the user.

	fmt.Println(color.Colorize(color.Red, "[*] This tool is for training."))
	fmt.Println(color.Colorize(color.Red, "[*]Enter checkwpjson -h to show help"))
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "domain",
				Value:       "",
				Aliases:     []string{"d"},
				Usage:       "Enter just one domain",
				Destination: &domains,
			},
			&cli.StringFlag{
				Name:        "list",
				Value:       "",
				Aliases:     []string{"l"},
				Usage:       "Enter a list from text file",
				Destination: &inputs,
			},
			&cli.BoolFlag{
				Name:    "pipe",
				Aliases: []string{"p"},
				Usage:   "Enter just from pipe line",
				Count:   &pipel,
			},

			&cli.StringFlag{
				Name:        "output",
				Value:       "output.csv",
				Aliases:     []string{"o"},
				Usage:       "Enter output csv file name  ",
				Destination: &outputs,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if domains != "" {
				results = withname(domains)

				writeResults(results, outputs)
			} else if inputs != "" {
				results = withlist(inputs)
				writeResults(results, outputs)
			} else if pipel > 0 {
				results = withpip()
				writeResults(results, outputs)
			}
			//	withlist("list", wg)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
