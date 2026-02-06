package ui

import (
	"fmt"
	"log"
)

func printSpinResult(spinResult [][]string) {
	for _, row := range spinResult {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

func getCommand() string {
	var command string
	fmt.Println("Please enter your bet or use one of commands: Quit, Restart")
	_, err := fmt.Scanln(&command)
	if err != nil {
		log.Fatal(err)
	}
	return command
}

func GetName() string {
	var name string
	fmt.Print("To start game, please enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func restartMessage() {
	fmt.Println("Your balance has been set to starting 250 chips")
}

func invalidBetMessage() {
	fmt.Println("You entered wrong bet, it must be positive number and must not be larger than balance, try again")
}

func printWinLose(profit uint, bet uint) {
	if profit == 0 {
		fmt.Printf("You lost $%d\n", bet)
	} else {
		fmt.Printf("You won $%d\n", profit-bet)
	}
}

func wrongCommandMessage(){
	fmt.Println("You entered wrong command, enter bet or one of commands QUIT or RESTART")
}

func printBalance(balance uint) {
	fmt.Printf("Your balance is $%d\n", balance)
}

func printGameOver(profit uint, flag bool) {
	if flag {
		fmt.Printf("We hope you have enjoyed your time, you have won: $%d\n", profit)
	} else {
		fmt.Printf("Today was a rough day for you, you lost: $%d\n", profit)
	}
}
