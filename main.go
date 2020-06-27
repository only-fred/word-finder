package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	answer := "YES"

	for answer == "YES" {
		hasBeenFound := true

		userWord := readUserWord()

		file, err := os.Open("./wordlist/wordlist.txt")
		if err != nil {
			log.Fatal()
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if line == userWord {
				fmt.Print("\n-------------------------------\n")
				fmt.Print("Your word has been found!")
				fmt.Print("\n-------------------------------\n")

				hasBeenFound = false

				break
			}
		}

		if hasBeenFound {
			fmt.Print("\n------------------------------------\n")
			fmt.Print("--Your word hasn't been found!--")
			fmt.Print("\n------------------------------------\n")
		}

		file.Close()

		fmt.Println("\nTry again? (YES/NO)")
		fmt.Scanf("%s", &answer)

		answer = strings.ToUpper(answer)
	}
}

func readUserWord() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nType a word ")
	fmt.Print("-> ")
	scanner.Scan()

	return scanner.Text()
}
