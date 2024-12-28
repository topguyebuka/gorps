package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("welcome to the Game of the Year")

	var playerScore int
	var computerScore int
	var numOfGamePlayed int
	var drawGames int
	playAgain := true
	var anotherGame string

	rand.Seed(time.Now().UnixNano())
	var answer string
	randomNum := rand.Intn(3) + 1
	var computer string

	for playAgain {
		answer = ""
		computer = ""
		fmt.Print("Choose Paper, Rock or Scissors:")
		fmt.Scan(&answer)
		answer = strings.ToUpper(answer)

		if answer != "ROCK" && answer != "PAPER" && answer != "SCISSORS" {
			fmt.Println("Invalid Response, Try Again")
			playAgain = true

		} else {

			fmt.Println("Player:", answer)

			switch randomNum {
			case 1:
				computer = "ROCK"
			case 2:
				computer = "PAPER"
			case 3:
				computer = "SCISSORS"
			}

			fmt.Println("Computer:", computer)

			switch answer {
			case "ROCK":
				if computer == "ROCK" {
					fmt.Println("You Draw")
					drawGames++

				} else if computer == "PAPER" {
					fmt.Println("You Lose")
					numOfGamePlayed++
					computerScore++

				} else {
					fmt.Println("You Win")
					playerScore++
					numOfGamePlayed++
				}
				break
			case "PAPER":
				if computer == "ROCK" {
					fmt.Println("You Win!")
					playerScore++
					numOfGamePlayed++

				} else if computer == "PAPER" {
					fmt.Println("You Draw!")
					numOfGamePlayed++
					drawGames++

				} else {
					fmt.Println("You Lose!")
					numOfGamePlayed++
					computerScore++
				}
				break
			case "SCISSORS":
				if computer == "ROCK" {
					fmt.Println("You Lose!")
					numOfGamePlayed++
					computerScore++

				} else if computer == "PAPER" {
					fmt.Println("You Win!")
					playerScore++
					numOfGamePlayed++

				} else {
					fmt.Println("You Draw!")
					numOfGamePlayed++
					drawGames++
				}
				break
			}
			fmt.Print("Do you want to play again?, select Y/N:")
			fmt.Scan(&anotherGame)
			anotherGame = strings.ToUpper(anotherGame)

			switch anotherGame {
			case "Y":
				playAgain = true
			case "N":
				fmt.Printf("you played %v games, you won %v games, computer won %v games and finally you draw %v games\n", numOfGamePlayed, playerScore, computerScore, drawGames)
				fmt.Println("Thank you for playing")
				playAgain = false
			}
		}

	}

}
