package main

import (
	"fmt"
	"sraynitjsr/behavioral"
	"sraynitjsr/creational"
	lld "sraynitjsr/low_level_design"
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
	lld.LLDUsingGoLang()
}
