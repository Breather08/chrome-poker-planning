package main

type Card struct {
	ID       string `json:"id"`
	Value    uint8  `json:"value"`
	IsPicked bool   `json:"isPicked"`
}
