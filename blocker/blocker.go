package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const (
	waitingTime = 20 * time.Minute
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

	} else if block == "false" {
		fmt.Printf("Befeore you start going down the rabbit hole, you should think about your future.\nDon't you want to be a master at programing?\nDon't you want to know how to draw?\nDon't you want to expand your horizons by reading books?\nDon't you think you have something better to do with your time than waste it on youtube?\nThink about your children and family, the sooner the actions are done, the greater the impact they have on the future.\n\n")
		time.Sleep(waitingTime)
		fmt.Println("Do you want to continue with unblocking youtube? (yes/no)")
		unblockConfirmed := askForConfirmation()

		if unblockConfirmed {
			fmt.Printf("Remebre to write down every video you watched, and it is mandatory to take notes.\n Enjoy youtube responsibly.\n")
			removeBlockadeFromFile(hosts)
		} else {
			fmt.Println("Thank you for investing time into your future.")
		}
	}
}

func removeBlockadeFromFile(fileName *os.File) {
	var contents []string

	r := bufio.NewReader(fileName)
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

func askForConfirmation() bool {
	response := ""

	yesResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	noResponses := []string{"n", "N", "no", "No", "NO"}

	fmt.Scanln(&response)
	if containsString(yesResponses, response) {
		return true
	} else if containsString(noResponses, response) {
		return false
	} else {
		fmt.Printf("Please answer with yes or no, and then hit Enter\n")
		return askForConfirmation()
	}
}

func containsString(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}

	return false
}
