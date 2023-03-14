package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func getNumberOfToken(walletAddress string) (int, error) {
	endpoint := fmt.Sprintf(
		"%s/api?module=account&action=tokenbalance&contractaddress=%s&address=%s&tag=latest&apikey=%s",
		os.Getenv("API_URL"),
		os.Getenv("CONTRACT_ADDRESS"),
		walletAddress,
		os.Getenv("API_TOKEN"),
	)

	response, err := http.Get(endpoint)

	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	var data TokenBalanceResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		return 0, err
	}

	numberOfTokens, _ := strconv.Atoi(data.Result)

	return numberOfTokens, nil
}

func getTokenSupply() (float64, error) {
	endpoint := fmt.Sprintf(
		"%s/api?module=stats&action=tokensupply&contractaddress=%s&apikey=%s",
		os.Getenv("API_URL"),
		os.Getenv("CONTRACT_ADDRESS"),
		os.Getenv("API_TOKEN"),
	)

	response, err := http.Get(endpoint)

	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	var data TokenSupplyResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		return 0, err
	}

	supply, _ := strconv.ParseFloat(data.Result, 64)

	return supply, nil
}