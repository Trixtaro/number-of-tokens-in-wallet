package main

type TokenBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type TokenSupplyResponse struct {
	TokenBalanceResponse
}