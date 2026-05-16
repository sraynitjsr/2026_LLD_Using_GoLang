# Abstract Factory Design Pattern

## Overview
The Abstract Factory pattern provides an interface for creating families of related or dependent objects without specifying their concrete classes. It's a creational pattern that abstracts the process of object creation.

## Intent
- Provide an interface for creating families of related or dependent objects
- Encapsulate a group of individual factories with a common theme
- Isolate concrete classes from the client
- Promote consistency among products

## When to Use
- System should be independent of how its products are created
- System should work with multiple families of products
- Family of related products is designed to be used together
- You want to provide a class library of products, revealing only their interfaces

## Structure

```
┌─────────────────┐
│  GUIFactory     │ (Abstract Factory)
├─────────────────┤
│ +CreateButton() │
│ +CreateCheckbox()│
└─────────────────┘
        △
        │
    ┌───┴───┐
    │       │
┌───┴────┐  ┌────┴────┐
│  Light │  │  Dark   │ (Concrete Factories)
│ Theme  │  │  Theme  │
│Factory │  │ Factory │
└────────┘  └─────────┘
```

## Example Implementation

```go
// Abstract Products
type Button interface {
    Render() string
}

type Checkbox interface {
    Render() string
}

// Abstract Factory
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}

// Concrete Products - Light Theme
type LightButton struct{}
func (b *LightButton) Render() string {
    return "Light Button"
}

type LightCheckbox struct{}
func (c *LightCheckbox) Render() string {
    return "Light Checkbox"
}

// Concrete Factory - Light Theme
type LightThemeFactory struct{}
func (f *LightThemeFactory) CreateButton() Button {
    return &LightButton{}
}
func (f *LightThemeFactory) CreateCheckbox() Checkbox {
    return &LightCheckbox{}
}

// Client Code
func RenderUI(factory GUIFactory) {
    button := factory.CreateButton()
    checkbox := factory.CreateCheckbox()
    
    fmt.Println(button.Render())
    fmt.Println(checkbox.Render())
}

// Usage
lightFactory := &LightThemeFactory{}
RenderUI(lightFactory)
```

## Real-World Examples

### 1. UI Theming
```go
// Different themes create consistent families of UI components
lightFactory := &LightThemeFactory{}
darkFactory := &DarkThemeFactory{}
highContrastFactory := &HighContrastFactory{}
```

### 2. Database Drivers
```go
// Different database vendors with consistent connection/command objects
mysqlFactory := &MySQLFactory{}
postgresFactory := &PostgreSQLFactory{}

connection := factory.CreateConnection()
command := factory.CreateCommand()
```

### 3. Document Converters
```go
// Different document formats with consistent readers/writers
pdfFactory := &PDFFactory{}
docxFactory := &DOCXFactory{}

reader := factory.CreateReader()
writer := factory.CreateWriter()
```

## Advantages

1. **Isolation of Concrete Classes**
   - Client code works with interfaces, not concrete implementations
   - Easy to change product families

2. **Consistency Among Products**
   - Ensures products from same family work together
   - Prevents mixing incompatible products

3. **Single Responsibility Principle**
   - Product creation code is isolated in one place

4. **Open/Closed Principle**
   - Easy to introduce new product families without changing existing code

5. **Type Safety**
   - Compiler ensures products from same family are used together

## Disadvantages

1. **Complexity**
   - Introduces many new interfaces and classes
   - Can be overkill for simple scenarios

2. **Difficult to Support New Product Types**
   - Adding new product to family requires changing all factories

3. **Code Overhead**
   - More code to write and maintain

## Abstract Factory vs Factory Method

| Aspect | Abstract Factory | Factory Method |
|--------|-----------------|----------------|
| Purpose | Creates families of related objects | Creates one type of object |
| Structure | Multiple factory methods | Single factory method |
| Complexity | More complex | Simpler |
| Use Case | Multiple related products | Single product variations |

## Best Practices

1. **Start Simple**
   ```go
   // Start with Factory Method, upgrade to Abstract Factory when needed
   ```

2. **Use Dependency Injection**
   ```go
   func NewApplication(factory GUIFactory) *Application {
       return &Application{factory: factory}
   }
   ```

3. **Configuration-Based Factory Selection**
   ```go
   func GetFactory(theme string) GUIFactory {
       switch theme {
       case "light":
           return &LightThemeFactory{}
       case "dark":
           return &DarkThemeFactory{}
       default:
           return &LightThemeFactory{}
       }
   }
   ```

4. **Combine with Singleton**
   ```go
   // Factories themselves can be singletons
   var lightFactory *LightThemeFactory
   var once sync.Once
   
   func GetLightFactory() *LightThemeFactory {
       once.Do(func() {
           lightFactory = &LightThemeFactory{}
       })
       return lightFactory
   }
   ```

## Common Pitfalls

1. **Overuse**: Don't use when simple factory or factory method suffices
2. **Rigid Design**: Hard to add new product types to families
3. **Too Many Classes**: Can lead to class explosion

## Testing Strategy

```go
// Mock factory for testing
type MockFactory struct{}

func (f *MockFactory) CreateButton() Button {
    return &MockButton{}
}

func (f *MockFactory) CreateCheckbox() Checkbox {
    return &MockCheckbox{}
}

// Test with mock
func TestRenderUI(t *testing.T) {
    mockFactory := &MockFactory{}
    RenderUI(mockFactory)
    // Assert expectations
}
```

## Related Patterns

- **Factory Method**: Abstract Factory uses factory methods to create products
- **Singleton**: Concrete factories are often singletons
- **Prototype**: Can use prototype for product creation instead of factory methods
- **Builder**: Similar abstraction level but focuses on construction steps

## Go-Specific Considerations

1. **Interface Segregation**: Keep interfaces small and focused
2. **Struct Embedding**: Can use embedding for partial implementations
3. **Function Types**: Can use function types as lightweight factories
4. **Generics**: Go 1.18+ generics can simplify some patterns

```go
// Using function types as factories
type ButtonFactory func() Button
type CheckboxFactory func() Checkbox

var lightButtonFactory ButtonFactory = func() Button {
    return &LightButton{}
}
```

## Conclusion

Abstract Factory is powerful for managing families of related objects. Use it when you need to ensure consistency among products and want to isolate product creation logic. However, be mindful of the added complexity and use it judiciously.
