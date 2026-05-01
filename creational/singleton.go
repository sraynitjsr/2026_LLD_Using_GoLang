package creational

import (
	"fmt"
	"sync"
)

// singleton is the private struct that ensures only one instance exists
type singleton struct {
	data string
}

var (
	instance *singleton // The single instance
	once     sync.Once  // Ensures thread-safe one-time initialization
)

// GetInstance returns the singleton instance (creates it on first call)
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "I'm A Singleton Data"}
	})
	return instance
}

// MySingleton demonstrates the singleton pattern
func MySingleton() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Printf("s1 == s2 => %v\n", s1 == s2)
	fmt.Printf("s1 != s2 => %v\n", s1 != s2)
	fmt.Printf("Singleton Data => %s\n", s1.data)
	fmt.Println("\nSingleton: single instance, thread-safe, lazy-init")
}
