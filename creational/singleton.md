# Singleton Design Pattern

## Overview
The Singleton pattern is a creational design pattern that ensures a class has only one instance throughout the application lifecycle and provides a global point of access to that instance.

## Intent
- Ensure a class has only one instance
- Provide a global access point to that instance
- Lazy initialization - instance created only when needed
- Thread-safe implementation to handle concurrent access

## When to Use
- **Configuration Management**: Single configuration object shared across the application
- **Logging**: One logger instance managing all log operations
- **Database Connections**: Single connection pool manager
- **Cache Management**: One cache instance for the entire application
- **Thread Pools**: Single thread pool coordinator
- **Device Drivers**: Single instance managing hardware access

## Minimal Implementation in Go

This implementation demonstrates the core singleton pattern in its simplest, most maintainable form:

```go
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
	fmt.Println("\n✓ Singleton: single instance, thread-safe, lazy-init")
}
```

## Key Principles

**Why `sync.Once`?**
- Guarantees the initialization function runs exactly once
- Thread-safe without explicit locks
- More efficient than mutex-based implementations
- Idiomatic Go approach

**Private Struct**
- Lowercase `singleton` name prevents external instantiation
- Only `GetInstance()` can create and return the instance

**Lazy Initialization**
- Instance is created only when first accessed
- No memory wasted if never used

## Usage Example

```go
// First call - creates the instance
s1 := GetInstance()

// Second call - returns the same instance
s2 := GetInstance()

// Verify both references point to the same object
fmt.Printf("s1 == s2 => %v\n", s1 == s2)     // Output: true
fmt.Printf("s1 != s2 => %v\n", s1 != s2)     // Output: false
fmt.Printf("Singleton Data => %s\n", s1.data) // Access the data
```

## Thread Safety

The singleton pattern handles concurrent access safely:

```go
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        inst := GetInstance() // All goroutines get the same instance
    }()
}
wg.Wait()
```

Even with 100 goroutines calling `GetInstance()` concurrently, only one instance is created.

## Advantages

✅ **Controlled Access**: Single point of access to the instance  
✅ **Lazy Initialization**: Instance created only when needed  
✅ **Thread-Safe**: Using `sync.Once` ensures safe concurrent access  
✅ **Global State**: Shared state across the application  
✅ **Memory Efficient**: Only one instance exists  

## Disadvantages

❌ **Global State**: Can make testing difficult  
❌ **Hidden Dependencies**: Classes depending on singleton have hidden coupling  
❌ **Violates Single Responsibility**: Class controls its instantiation and business logic  
❌ **Difficult to Mock**: Hard to replace with test doubles  
❌ **Concurrency Issues**: Shared mutable state requires careful synchronization  

## Best Practices

1. **Use `sync.Once`**: Idiomatic Go approach for thread-safe singleton
2. **Private Struct**: Lowercase struct name prevents external instantiation
3. **Immutable Where Possible**: Minimize mutable state to reduce concurrency issues
4. **Dependency Injection**: Consider alternatives like dependency injection for testability
5. **Interface-Based**: Return interfaces instead of concrete types for flexibility

## Alternatives to Consider

- **Dependency Injection**: Pass dependencies explicitly rather than using global state
- **Context Pattern**: Use `context.Context` to pass request-scoped data
- **Package-Level Functions**: Simple package-level functions for stateless operations
- **Interface-Based Design**: Define interfaces and inject implementations

## Real-World Go Examples

Standard library patterns:
```go
import (
    "log"   // Package-level default logger
    "rand"  // Global random number generator
)
```

Custom singleton examples:
```go
// Configuration manager
var (
    config *Config
    configOnce sync.Once
)

func GetConfig() *Config {
    configOnce.Do(func() {
        config = &Config{/* load from file */}
    })
    return config
}

// Database connection pool
var (
    dbPool *sql.DB
    dbOnce sync.Once
)

func GetDB() *sql.DB {
    dbOnce.Do(func() {
        dbPool, _ = sql.Open("postgres", "...")
    })
    return dbPool
}
```

## Related Patterns

- **Factory Method**: Can use singleton factory to create objects
- **Abstract Factory**: Factory itself can be a singleton
- **Builder**: Builder instance can be a singleton
- **Facade**: Facade often implemented as singleton

## Conclusion

The Singleton pattern in Go is elegantly simple: a private struct, package-level variables, and `sync.Once` for thread-safe initialization. This minimal implementation (~30 lines) provides:

- **Single instance guarantee**
- **Thread-safe lazy initialization**
- **Clean, idiomatic Go code**

While powerful, use singletons judiciously. Consider dependency injection for better testability and flexibility when appropriate. The pattern works best for truly global resources like configuration managers, loggers, and connection pools.
