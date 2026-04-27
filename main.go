package main

import (
	"fmt"
	"sraynitjsr/behavioral"
	"sraynitjsr/creational"
	"sraynitjsr/structural"
)

func main() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║   Full Fledge Low Level System Design Using GoLang            ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	behavioral.StartBehavioral()

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	creational.StartCreational()

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	structural.StartStructural()
}
