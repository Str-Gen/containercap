package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

func main() {
	log.Println("Welcome to containercap!")
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
}
