package creational

import "fmt"

// Builder Pattern - Constructs complex objects step by step

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

// Display shows the computer configuration
func (c *Computer) Display() {
	fmt.Println("    ┌─────────────────────────────────────────┐")
	fmt.Printf("    │ CPU:      %-30s│\n", c.CPU)
	fmt.Printf("    │ RAM:      %-30s│\n", fmt.Sprintf("%d GB", c.RAM))
	fmt.Printf("    │ Storage:  %-30s│\n", fmt.Sprintf("%d GB", c.Storage))
	fmt.Printf("    │ GPU:      %-30s│\n", c.GPU)
	fmt.Printf("    │ OS:       %-30s│\n", c.OS)
	fmt.Printf("    │ Monitor:  %-30s│\n", c.Monitor)
	fmt.Printf("    │ Keyboard: %-30s│\n", c.Keyboard)
	fmt.Println("    └─────────────────────────────────────────┘")
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
	return b
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

// Director knows how to build specific configurations
type Director struct {
	builder ComputerBuilder
}

// NewDirector creates a new director with a builder
func NewDirector(builder ComputerBuilder) *Director {
	return &Director{builder: builder}
}

// BuildGamingComputer constructs a gaming computer
func (d *Director) BuildGamingComputer() *Computer {
	return d.builder.
		SetCPU("Intel i9-13900K").
		SetRAM(32).
		SetStorage(2000).
		SetGPU("NVIDIA RTX 4090").
		SetOS("Windows 11 Pro").
		SetMonitor("32 inch 4K Gaming").
		SetKeyboard("Mechanical RGB").
		Build()
}

// BuildOfficeComputer constructs an office computer
func (d *Director) BuildOfficeComputer() *Computer {
	return d.builder.
		SetCPU("Intel i5-12400").
		SetRAM(16).
		SetStorage(512).
		SetGPU("Integrated Graphics").
		SetOS("Windows 11").
		SetMonitor("24 inch Full HD").
		SetKeyboard("Wireless Keyboard").
		Build()
}

// MyBuilder demonstrates the builder pattern
func MyBuilder() {
	fmt.Println("    Building Gaming Computer:")
	gamingBuilder := NewDesktopBuilder()
	director := NewDirector(gamingBuilder)
	gamingPC := director.BuildGamingComputer()
	gamingPC.Display()

	fmt.Println()
	fmt.Println("    Building Office Computer:")
	officeBuilder := NewDesktopBuilder()
	director = NewDirector(officeBuilder)
	officePC := director.BuildOfficeComputer()
	officePC.Display()

	fmt.Println()
	fmt.Println("    Building Custom Computer (without Director):")
	customPC := NewDesktopBuilder().
		SetCPU("AMD Ryzen 9 7950X").
		SetRAM(64).
		SetStorage(4000).
		SetGPU("AMD Radeon RX 7900 XTX").
		SetOS("Linux Ubuntu 22.04").
		SetMonitor("34 inch Ultrawide").
		SetKeyboard("Ergonomic Split").
		Build()
	customPC.Display()
}
