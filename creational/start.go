package creational

import "fmt"

func StartCreational() {
	fmt.Println("Creational Design Patterns")
	fmt.Println()

	fmt.Println("Singleton Design Pattern => Single Instance, Thread-Safe, Lazy-Initialization")
	MySingleton()
}
