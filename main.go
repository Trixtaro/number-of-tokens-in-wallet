package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if len(os.Args) == 1 {
		fmt.Println("You need to introduce the wallet address as first argument.")
		return
	}

	numberOfTokens, err := getNumberOfToken(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("Total number of tokens: %d", numberOfTokens))
	if (numberOfTokens > 0) {
		// Mumbay PolygonScan doesn't provide a lot of information about tokens, so I decided to show token supply if the
		// wallet address balance is greater than 0
		supply, err := getTokenSupply()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(fmt.Sprintf("Token supply: %.2f", supply))
	}
}
