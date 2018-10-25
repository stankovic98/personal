package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	block := os.Getenv("BLOCK")
	if block == "" {
		log.Printf("Block variable not set")
	}

	hosts, err := os.OpenFile("hosts", os.O_APPEND|os.O_RDWR, 0644)
	defer hosts.Close()
	if err != nil {
		log.Printf("err at opening the file: %s\n", err)
	}

	if block == "true" {
		sites, err := ioutil.ReadFile("blockedSites")
		if err != nil {
			log.Print(err)
		}

		sites = append([]byte("\n"), sites...)
		_, err = hosts.Write(sites)
		if err != nil {
			log.Print(err)
		}

	} else {
		var contents []string

		r := bufio.NewReader(hosts)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {

			line := scanner.Text()
			contents = append(contents, line)

			if strings.Contains(scanner.Text(), "###") {

				forFile := []byte(strings.Join(contents, "\n"))
				ioutil.WriteFile("hosts", forFile, 0666)
				break
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
}
