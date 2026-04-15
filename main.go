package main

import (
	"fmt"
	"sync"
)

// Config represents application configuration (singleton)
type Config struct {
	AppName string
	Version string
	Port    int
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton instance
// ✅ Thread-safe (sync.Once)
// ✅ Lazy initialization (created only when first called)
// ✅ Zero overhead (atomic operations after first call)
// ✅ Memory efficient (single instance)
func GetConfig() *Config {
	once.Do(func() {
		fmt.Println("Creating singleton instance...")
		instance = &Config{
			AppName: "MyApp",
			Version: "1.0.0",
			Port:    8080,
		}
	})
	return instance
}

func main() {
	// First call - creates the instance
	config1 := GetConfig()
	fmt.Printf("Config 1: %s v%s (Port: %d)\n",
		config1.AppName, config1.Version, config1.Port)

	// Second call - returns existing instance
	config2 := GetConfig()
	fmt.Printf("Config 2: %s v%s (Port: %d)\n",
		config2.AppName, config2.Version, config2.Port)

	// Multiple goroutines accessing same singleton
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := GetConfig()
			fmt.Printf("Goroutine %d: %s\n", id, cfg.AppName)
		}(i)
	}
	wg.Wait()
}
