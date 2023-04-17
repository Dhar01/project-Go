// Guessing game : challenging players to guess a random number.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// With the help of the book "Head-first Go", this project was built.

func main() {

	// generating a random number
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// get user input & process & finalize the result
	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		// take input
		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// processing the input
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		// logic
		if guess < target {
			fmt.Println("Opps! Your guess was low.")
		} else if guess > target {
			fmt.Println("Your guess is high.")
		} else if guess == target {
			fmt.Println("Your guess is right! It is", target)
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

	fmt.Println("It was a good game!")
}
