package creational

import "fmt"

// Abstract Factory Pattern - Creates families of related objects without specifying their concrete classes

// Button interface - Abstract Product A
type Button interface {
	Render() string
}

// Checkbox interface - Abstract Product B
type Checkbox interface {
	Render() string
}

// GUIFactory interface - Abstract Factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Light Theme Components
type LightButton struct{}

func (b *LightButton) Render() string {
	return "Light Button [  Click  ]"
}

type LightCheckbox struct{}

func (c *LightCheckbox) Render() string {
	return "Light Checkbox [ ☐ ]"
}

// Dark Theme Components
type DarkButton struct{}

func (b *DarkButton) Render() string {
	return "Dark Button [  CLICK  ]"
}

type DarkCheckbox struct{}

func (c *DarkCheckbox) Render() string {
	return "Dark Checkbox [ ▢ ]"
}

// Light Theme Factory
type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button {
	return &LightButton{}
}

func (f *LightThemeFactory) CreateCheckbox() Checkbox {
	return &LightCheckbox{}
}

// Dark Theme Factory
type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button {
	return &DarkButton{}
}

func (f *DarkThemeFactory) CreateCheckbox() Checkbox {
	return &DarkCheckbox{}
}

// Client code that works with factories and products through abstract interfaces
func RenderUI(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	fmt.Println("    " + button.Render())
	fmt.Println("    " + checkbox.Render())
}

func MyAbstractFactory() {
	fmt.Println("    Creating Light Theme UI:")
	lightFactory := &LightThemeFactory{}
	RenderUI(lightFactory)

	fmt.Println()
	fmt.Println("    Creating Dark Theme UI:")
	darkFactory := &DarkThemeFactory{}
	RenderUI(darkFactory)
}
