package main

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{AppName: "MyApp"}
	})
	return instance
}

func main() {
	singletonExample()
}

func singletonExample() {
	config1 := GetConfig()
	config2 := GetConfig()

	fmt.Println("config1:", config1.AppName)
	fmt.Println("config2:", config2.AppName)

	fmt.Println("Same instance?", config1 == config2) // true
	fmt.Println("Same instance?", config1 != config2) // false
}
