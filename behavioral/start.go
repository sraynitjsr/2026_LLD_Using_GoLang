package behavioral

import "fmt"

func StartBehavioral() {
	cyan := "\033[1;36m"
	reset := "\033[0m"

	fmt.Println(cyan + "Behavioral Design Patterns" + reset)
	fmt.Println()

	fmt.Println(cyan + "[1] CHAIN OF RESPONSIBILITY DESIGN PATTERN" + reset)
	fmt.Println("    Description: Pass Request Along Handler Chain")
	fmt.Println()
	MyChainOfResponsibility()
	fmt.Println()
}
