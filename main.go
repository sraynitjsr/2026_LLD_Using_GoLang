package main

import (
	"fmt"
	"sraynitjsr/behavioral"
	"sraynitjsr/creational"
	lldinterview "sraynitjsr/lld_interview"
	"sraynitjsr/structural"
)

func main() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║   Full Fledge Low Level System Design Using Latest GoLang     ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	creational.StartCreational()

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	behavioral.StartBehavioral()

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	structural.StartStructural()

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	lldinterview.ShortenURL()
}
