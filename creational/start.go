package creational

import "fmt"

func StartCreational() {
	blue := "\033[1;34m"
	reset := "\033[0m"

	fmt.Println(blue + "Creational Design Patterns" + reset)
	fmt.Println()

	fmt.Println(blue + "[1] SINGLETON DESIGN PATTERN" + reset)
	fmt.Println("    Description: Single Instance, Thread-Safe, Lazy-Initialization")
	fmt.Println()
	MySingleton()

	fmt.Println()
	fmt.Println(blue + "[2] FACTORY METHOD DESIGN PATTERN" + reset)
	fmt.Println("    Description: Define Interface For Creating Objects")
	fmt.Println()
	MyFactoryMethod()

	fmt.Println()
	fmt.Println(blue + "[3] ABSTRACT FACTORY DESIGN PATTERN" + reset)
	fmt.Println("    Description: Create Families Of Related Objects")
	fmt.Println()
	MyAbstractFactory()
	fmt.Println()
}
