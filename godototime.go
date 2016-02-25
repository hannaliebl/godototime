package main

import (
	"fmt"
	"os"
)

const steamHost = "http://api.steampowered.com"

func main() {
	steamKey := os.Getenv("STEAM_KEY")

	fmt.Printf("steam key is: %v\n", steamKey)
}
