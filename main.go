package main

import (
	"go-enigma-laundry/delivery"
)

func main() {
	delivery.NewServer().Run()
}