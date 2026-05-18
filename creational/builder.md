# Builder Design Pattern

## Overview
The Builder pattern is a creational design pattern that separates the construction of a complex object from its representation, allowing the same construction process to create different representations. It constructs complex objects step by step and allows you to produce different types and representations using the same construction code.

## Intent
- Separate construction of a complex object from its representation
- Allow step-by-step object construction
- Provide control over the construction process
- Create different representations using the same construction process
- Avoid "telescoping constructor" anti-pattern

## When to Use
- **Complex Object Creation**: When creating objects with many optional parameters or configurations
- **Multiple Representations**: When you need to create different representations of the same object
- **Step-by-Step Construction**: When object creation requires multiple steps or initialization logic
- **Immutable Objects**: When building immutable objects that need all properties set before use
- **Readability**: When constructor has too many parameters making code hard to read
- **Validation**: When you need to validate object state before creation

## Real-World Use Cases
- **Computer Configuration**: Building PCs with various components (CPU, RAM, GPU, etc.)
- **Document Generation**: Creating documents with different sections and formats
- **HTTP Request Building**: Constructing complex HTTP requests with headers, params, body
- **SQL Query Construction**: Building complex SQL queries step by step
- **UI Components**: Creating complex UI elements with many optional properties
- **Configuration Objects**: Building application configuration from various sources
- **Test Data Creation**: Creating test objects with different property combinations

## Implementation in Go

This implementation demonstrates building computer configurations:

```go
package creational

import "fmt"

// Computer is the complex product we want to build
type Computer struct {
	CPU      string
	RAM      int
	Storage  int
	GPU      string
	OS       string
	Monitor  string
	Keyboard string
}

// ComputerBuilder interface defines the building steps
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram int) ComputerBuilder
	SetStorage(storage int) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	SetOS(os string) ComputerBuilder
	SetMonitor(monitor string) ComputerBuilder
	SetKeyboard(keyboard string) ComputerBuilder
	Build() *Computer
}

// DesktopBuilder is a concrete builder for desktop computers
type DesktopBuilder struct {
	computer *Computer
}

// NewDesktopBuilder creates a new desktop builder with defaults
func NewDesktopBuilder() *DesktopBuilder {
	return &DesktopBuilder{
		computer: &Computer{
			CPU:      "Intel i5",
			RAM:      8,
			Storage:  256,
			GPU:      "Integrated",
			OS:       "Windows 11",
			Monitor:  "24 inch LCD",
			Keyboard: "Standard",
		},
	}
}

func (b *DesktopBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b // Return self for method chaining
}

func (b *DesktopBuilder) SetRAM(ram int) ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *DesktopBuilder) SetStorage(storage int) ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *DesktopBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *DesktopBuilder) SetOS(os string) ComputerBuilder {
	b.computer.OS = os
	return b
}

func (b *DesktopBuilder) SetMonitor(monitor string) ComputerBuilder {
	b.computer.Monitor = monitor
	return b
}

func (b *DesktopBuilder) SetKeyboard(keyboard string) ComputerBuilder {
	b.computer.Keyboard = keyboard
	return b
}

func (b *DesktopBuilder) Build() *Computer {
	return b.computer
}
```

## Key Components

### 1. Product (Computer)
The complex object being constructed. Contains multiple fields that need to be configured.

### 2. Builder Interface (ComputerBuilder)
Declares the construction steps that are common to all builders. Each method returns the builder itself for method chaining (fluent interface).

### 3. Concrete Builder (DesktopBuilder)
Implements the builder interface and maintains the product being built. Provides default values and allows step-by-step configuration.

### 4. Director (Optional)
Knows how to use the builder to construct specific configurations. Encapsulates common building recipes.

```go
// Director knows how to build specific configurations
type Director struct {
	builder ComputerBuilder
}

func (d *Director) BuildGamingComputer() *Computer {
	return d.builder.
		SetCPU("Intel i9-13900K").
		SetRAM(32).
		SetStorage(2000).
		SetGPU("NVIDIA RTX 4090").
		Build()
}

func (d *Director) BuildOfficeComputer() *Computer {
	return d.builder.
		SetCPU("Intel i5-12400").
		SetRAM(16).
		SetStorage(512).
		Build()
}
```

## Usage Examples

### Without Director (Direct Building)
```go
// Fluent interface allows readable, chainable configuration
customPC := NewDesktopBuilder().
	SetCPU("AMD Ryzen 9 7950X").
	SetRAM(64).
	SetStorage(4000).
	SetGPU("AMD Radeon RX 7900 XTX").
	SetOS("Linux Ubuntu 22.04").
	Build()
```

### With Director (Predefined Configurations)
```go
builder := NewDesktopBuilder()
director := NewDirector(builder)

// Use predefined recipes
gamingPC := director.BuildGamingComputer()
officePC := director.BuildOfficeComputer()
```

## Key Principles

### Method Chaining (Fluent Interface)
Each setter method returns the builder itself, allowing methods to be chained in a readable way:
```go
builder.SetCPU("Intel i9").SetRAM(32).SetGPU("RTX 4090").Build()
```

### Default Values
Provide sensible defaults in the constructor so not all fields need to be set:
```go
func NewDesktopBuilder() *DesktopBuilder {
	return &DesktopBuilder{
		computer: &Computer{
			CPU: "Intel i5",  // Default value
			RAM: 8,           // Default value
			// ...
		},
	}
}
```

### Separation of Concerns
- **Builder**: Knows how to construct the object
- **Director**: Knows which steps to execute in which order
- **Product**: Knows nothing about how it's built

## Benefits

1. **Readability**: Code is much more readable than constructors with many parameters
```go
// Bad: Too many parameters
computer := NewComputer("Intel i9", 32, 2000, "RTX 4090", "Windows 11", "32in", "RGB")

// Good: Builder pattern
computer := NewDesktopBuilder().
	SetCPU("Intel i9").
	SetRAM(32).
	SetStorage(2000).
	Build()
```

2. **Flexibility**: Can construct objects with different configurations using the same builder

3. **Control**: Fine-grained control over the construction process

4. **Reusability**: Director can reuse the same builder to create different products

5. **Immutability**: Can build immutable objects by setting all properties before final construction

6. **Single Responsibility Principle**: Construction code is separated from business logic

7. **Validation**: Can validate the object state in the Build() method before returning

## Drawbacks

1. **Complexity**: Increases the number of classes in the codebase
2. **Overhead**: May be overkill for simple objects with few properties
3. **Mutability**: The builder itself is mutable during construction

## Comparison with Other Patterns

### Builder vs Factory Method
- **Builder**: Constructs complex objects step by step
- **Factory Method**: Creates objects in one step

### Builder vs Abstract Factory
- **Builder**: Focuses on constructing a single complex object step by step
- **Abstract Factory**: Creates families of related objects

### Builder vs Prototype
- **Builder**: Constructs objects from scratch step by step
- **Prototype**: Creates objects by copying existing ones

## Go-Specific Considerations

### Functional Options Pattern
Go has an alternative idiomatic pattern for optional parameters:

```go
type Option func(*Computer)

func WithCPU(cpu string) Option {
	return func(c *Computer) {
		c.CPU = cpu
	}
}

func NewComputer(opts ...Option) *Computer {
	c := &Computer{/* defaults */}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Usage
pc := NewComputer(
	WithCPU("Intel i9"),
	WithRAM(32),
	WithGPU("RTX 4090"),
)
```

Both patterns are valid in Go. Use:
- **Builder**: When you need multiple steps, validation, or the Director pattern
- **Functional Options**: When you primarily need optional parameters with defaults

## Best Practices

1. **Return Builder from Setters**: Enable method chaining for fluent interface
2. **Provide Defaults**: Set reasonable default values in the constructor
3. **Validate in Build()**: Check object validity before returning the product
4. **Use Director**: For common configurations that are reused frequently
5. **Make Builder Package-Private**: If the builder is only used within the package
6. **Consider Immutability**: Make the product immutable after Build() is called
7. **Document Required Fields**: Clearly indicate which fields must be set

## Testing Advantages

The Builder pattern makes testing easier:

```go
func TestComputerBuilder(t *testing.T) {
	// Easy to create test objects with specific configurations
	testPC := NewDesktopBuilder().
		SetCPU("Test CPU").
		SetRAM(16).
		Build()
	
	assert.Equal(t, "Test CPU", testPC.CPU)
	assert.Equal(t, 16, testPC.RAM)
}

// Create different test scenarios easily
gamingPC := director.BuildGamingComputer()
officePC := director.BuildOfficeComputer()
```

## Conclusion

The Builder pattern is excellent for constructing complex objects with many optional parameters. It improves code readability, provides flexibility in object creation, and separates construction logic from business logic. While it adds some complexity, the benefits in readability and maintainability often outweigh the cost for complex object creation scenarios.
