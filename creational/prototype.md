# Prototype Design Pattern

## Overview
The Prototype pattern is a creational design pattern that creates new objects by copying existing objects (prototypes) rather than creating new instances from scratch. It allows you to clone objects without coupling to their specific classes and can be more efficient when object creation is expensive.

## Intent
- Create new objects by cloning existing prototypes
- Avoid expensive initialization/creation operations
- Reduce subclassing needed for object creation
- Hide complexity of creating new instances
- Support adding and removing objects at runtime

## When to Use
- **Expensive Object Creation**: When creating new objects is resource-intensive (database queries, network calls, complex initialization)
- **Similar Objects**: When you need many similar objects with minor variations
- **Runtime Configuration**: When object types are determined at runtime
- **Avoiding Subclasses**: When you want to avoid creating a factory hierarchy
- **State Preservation**: When you need to preserve the current state of an object
- **Undo/Redo Operations**: When implementing undo/redo functionality

## Real-World Use Cases
- **Document Templates**: Creating new documents from predefined templates
- **Game Development**: Cloning game objects (enemies, items, terrain)
- **Configuration Management**: Creating configurations from templates
- **UI Components**: Cloning UI widgets with preset configurations
- **Database Records**: Creating new records based on existing ones
- **Graphics Editors**: Copying and pasting shapes, images, or layers
- **Caching**: Creating cached object instances

## Implementation in Go

This implementation demonstrates document management using prototypes:

```go
package creational

import "fmt"

// Document interface defines the Clone method
type Document interface {
	Clone() Document
	Display()
	SetTitle(title string)
	SetContent(content string)
}

// ReportDocument is a concrete prototype for reports
type ReportDocument struct {
	Title      string
	Content    string
	Author     string
	Template   string
	ReportType string
	Charts     []string
}

// Clone creates a deep copy of the report document
func (d *ReportDocument) Clone() Document {
	// Deep copy - create new slice for Charts
	charts := make([]string, len(d.Charts))
	copy(charts, d.Charts)

	return &ReportDocument{
		Title:      d.Title,
		Content:    d.Content,
		Author:     d.Author,
		Template:   d.Template,
		ReportType: d.ReportType,
		Charts:     charts,
	}
}

func (d *ReportDocument) SetTitle(title string) {
	d.Title = title
}

func (d *ReportDocument) SetContent(content string) {
	d.Content = content
}

// InvoiceDocument is a concrete prototype for invoices
type InvoiceDocument struct {
	Title         string
	Content       string
	Author        string
	Template      string
	InvoiceNumber string
	TotalAmount   float64
	Currency      string
}

// Clone creates a deep copy of the invoice document
func (d *InvoiceDocument) Clone() Document {
	return &InvoiceDocument{
		Title:         d.Title,
		Content:       d.Content,
		Author:        d.Author,
		Template:      d.Template,
		InvoiceNumber: d.InvoiceNumber,
		TotalAmount:   d.TotalAmount,
		Currency:      d.Currency,
	}
}

func (d *InvoiceDocument) SetTitle(title string) {
	d.Title = title
}

func (d *InvoiceDocument) SetContent(content string) {
	d.Content = content
}
```

## Key Components

### 1. Prototype Interface (Document)
Declares the `Clone()` method that all concrete prototypes must implement.

### 2. Concrete Prototypes (ReportDocument, InvoiceDocument)
Implement the Clone method to create copies of themselves. Must handle deep copying of reference types.

### 3. Prototype Registry (Optional but Recommended)
Manages a collection of prototype objects and provides methods to retrieve and clone them.

```go
// DocumentRegistry manages prototype instances
type DocumentRegistry struct {
	prototypes map[string]Document
}

func NewDocumentRegistry() *DocumentRegistry {
	return &DocumentRegistry{
		prototypes: make(map[string]Document),
	}
}

func (r *DocumentRegistry) Register(key string, prototype Document) {
	r.prototypes[key] = prototype
}

func (r *DocumentRegistry) Create(key string) Document {
	if prototype, ok := r.prototypes[key]; ok {
		return prototype.Clone()
	}
	return nil
}
```

## Deep Copy vs Shallow Copy

### Shallow Copy (❌ Avoid for Reference Types)
```go
// Bad: Shallow copy - both objects share the same Charts slice
func (d *ReportDocument) Clone() Document {
	return &ReportDocument{
		Charts: d.Charts, // ❌ Both point to same slice
	}
}
```

### Deep Copy (✅ Correct Approach)
```go
// Good: Deep copy - each object has its own Charts slice
func (d *ReportDocument) Clone() Document {
	charts := make([]string, len(d.Charts))
	copy(charts, d.Charts) // ✅ Create new slice
	
	return &ReportDocument{
		Charts: charts,
	}
}
```

## Usage Examples

### Basic Cloning
```go
// Create a prototype
original := &ReportDocument{
	Author:     "System",
	Template:   "Standard Report Template",
	ReportType: "Quarterly",
	Charts:     []string{"Sales", "Revenue"},
}

// Clone and customize
clone := original.Clone().(*ReportDocument)
clone.SetTitle("Q1 2026 Report")
clone.SetContent("Q1 analysis")
```

### Using Registry
```go
// Set up registry with prototypes
registry := NewDocumentRegistry()
registry.Register("report", reportPrototype)
registry.Register("invoice", invoicePrototype)

// Create new documents from prototypes
q1Report := registry.Create("report")
q1Report.SetTitle("Q1 Report")

q2Report := registry.Create("report")
q2Report.SetTitle("Q2 Report")

invoice := registry.Create("invoice")
invoice.SetTitle("Invoice #001")
```

## Key Principles

### 1. Cloneable Interface
All prototypes must implement a Clone method that returns a new instance:
```go
type Cloneable interface {
	Clone() Cloneable
}
```

### 2. Deep Copy for Reference Types
Always create deep copies of slices, maps, and pointers:
```go
// Deep copy a slice
newSlice := make([]string, len(original.Slice))
copy(newSlice, original.Slice)

// Deep copy a map
newMap := make(map[string]int)
for k, v := range original.Map {
	newMap[k] = v
}
```

### 3. Registry Pattern
Use a registry to manage and access prototypes:
- Centralized prototype management
- Runtime prototype registration
- Easy prototype retrieval by key

## Benefits

1. **Performance**: Cloning can be faster than creating from scratch when initialization is expensive
```go
// Expensive to create
original := CreateExpensiveObject() // Database query, complex initialization

// Fast to clone
clone1 := original.Clone()
clone2 := original.Clone()
```

2. **Reduced Complexity**: Avoid complex initialization logic for each new object

3. **Runtime Flexibility**: Add and remove prototypes at runtime

4. **Avoid Subclassing**: No need to create factory hierarchies for different object types

5. **State Preservation**: Clone objects with their current state intact

6. **Hide Complexity**: Client doesn't need to know how objects are created

## Drawbacks

1. **Deep Copy Complexity**: Implementing deep copy for complex objects with circular references can be challenging

2. **Clone Method Maintenance**: Must update Clone() method when adding new fields

3. **Type Assertions**: May need type assertions when working with interface types
```go
report := registry.Create("report").(*ReportDocument) // Type assertion needed
```

4. **Not Always Appropriate**: For simple objects, direct instantiation may be simpler

## Comparison with Other Patterns

### Prototype vs Factory Method
- **Prototype**: Clones existing objects
- **Factory Method**: Creates new objects from scratch

### Prototype vs Abstract Factory
- **Prototype**: Uses composition (cloning)
- **Abstract Factory**: Uses inheritance (factory hierarchy)

### Prototype vs Builder
- **Prototype**: Starts with a pre-configured object and clones it
- **Builder**: Constructs objects step by step from scratch

### Prototype vs Singleton
- **Prototype**: Creates multiple similar instances
- **Singleton**: Ensures only one instance exists

## Go-Specific Considerations

### Using copy() for Slices
Go provides built-in `copy()` function for slice deep copying:
```go
newSlice := make([]string, len(original))
copy(newSlice, original)
```

### Struct Literal Copying
Simple structs with value types are automatically copied:
```go
original := Point{X: 10, Y: 20}
clone := original // Automatic copy for value types
```

### Pointer vs Value Receivers
Be careful with pointer vs value receivers:
```go
// Returns pointer to new object
func (d *Document) Clone() *Document {
	return &Document{...}
}
```

### JSON-based Cloning (Alternative Approach)
For complex objects, you can use JSON serialization:
```go
func DeepCopy(src, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}
```
Note: This is slower but handles complex nested structures automatically.

## Best Practices

1. **Implement Deep Copy**: Always deep copy reference types (slices, maps, pointers)

2. **Use Registry**: Implement a registry for managing multiple prototype types

3. **Document Clone Behavior**: Clearly document whether Clone() performs deep or shallow copy

4. **Test Clone Independence**: Ensure modifying cloned objects doesn't affect originals

5. **Consider Immutability**: After cloning, consider making the clone immutable if appropriate

6. **Handle Circular References**: Be careful with objects that reference themselves or each other

7. **Benchmark**: Measure if cloning is actually faster than creation for your use case

## Testing the Prototype Pattern

```go
func TestPrototypeCloning(t *testing.T) {
	original := &ReportDocument{
		Title:  "Original",
		Charts: []string{"Chart1"},
	}
	
	clone := original.Clone().(*ReportDocument)
	clone.Title = "Clone"
	clone.Charts[0] = "Chart2"
	
	// Ensure independence
	assert.Equal(t, "Original", original.Title)
	assert.Equal(t, "Chart1", original.Charts[0])
	assert.Equal(t, "Clone", clone.Title)
	assert.Equal(t, "Chart2", clone.Charts[0])
}
```

## Conclusion

The Prototype pattern is powerful when you need to create many similar objects or when object creation is expensive. It provides a flexible alternative to factory patterns and can significantly improve performance in scenarios where cloning is cheaper than creation. The key is implementing proper deep copying and using a registry to manage prototypes effectively.
