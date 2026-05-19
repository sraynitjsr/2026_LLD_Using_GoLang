# Chain of Responsibility Design Pattern

## Overview
The Chain of Responsibility pattern is a behavioral design pattern that allows passing requests along a chain of handlers. Each handler decides either to process the request or to pass it to the next handler in the chain.

## Intent
- Avoid coupling the sender of a request to its receiver
- Give multiple objects a chance to handle the request
- Chain the receiving objects and pass the request along the chain
- Each handler can process or forward the request independently

## When to Use
- **Request Processing Pipeline**: Multiple objects may handle a request, handler isn't known beforehand
- **Logging Systems**: Different log levels handled by different loggers
- **Authentication/Authorization**: Multiple security checks in sequence
- **Event Bubbling**: UI events propagating through component hierarchy
- **Support Ticket Systems**: Escalation based on complexity/priority
- **Validation Chains**: Multiple validation rules applied sequentially
- **Approval Workflows**: Document approval by multiple authorities

## Real-World Examples
- Exception handling in programming languages (try-catch bubbling)
- Servlet filters in Java web applications
- Middleware in Express.js/Node.js applications
- Event propagation in DOM (capture/bubble phases)
- Email spam filters (multiple filters in sequence)
- HTTP request processing in web frameworks

## Complete Implementation in Go

```go
package behavioral

import "fmt"

// SupportTicket represents a support request
type SupportTicket struct {
	Issue       string
	Priority    int // 1=Low, 2=Medium, 3=High
	Description string
}

// SupportHandler interface defines the handler contract
type SupportHandler interface {
	SetNext(handler SupportHandler) SupportHandler
	Handle(ticket *SupportTicket)
}

// BaseSupportHandler provides default chaining behavior
type BaseSupportHandler struct {
	next SupportHandler
}

func (h *BaseSupportHandler) SetNext(handler SupportHandler) SupportHandler {
	h.next = handler
	return handler
}

// Level1Support handles basic issues (Priority 1)
type Level1Support struct {
	BaseSupportHandler
}

func (h *Level1Support) Handle(ticket *SupportTicket) {
	if ticket.Priority == 1 {
		fmt.Printf("✓ [Level 1 Support] Handling: %s\n", ticket.Issue)
		fmt.Printf("  Resolution: %s\n", ticket.Description)
	} else if h.next != nil {
		fmt.Printf("→ [Level 1 Support] Escalating: %s\n", ticket.Issue)
		h.next.Handle(ticket)
	}
}

// Level2Support handles medium complexity issues (Priority 2)
type Level2Support struct {
	BaseSupportHandler
}

func (h *Level2Support) Handle(ticket *SupportTicket) {
	if ticket.Priority == 2 {
		fmt.Printf("✓ [Level 2 Support] Handling: %s\n", ticket.Issue)
		fmt.Printf("  Resolution: %s\n", ticket.Description)
	} else if h.next != nil {
		fmt.Printf("→ [Level 2 Support] Escalating: %s\n", ticket.Issue)
		h.next.Handle(ticket)
	}
}

// ManagerSupport handles high priority issues (Priority 3)
type ManagerSupport struct {
	BaseSupportHandler
}

func (h *ManagerSupport) Handle(ticket *SupportTicket) {
	if ticket.Priority == 3 {
		fmt.Printf("✓ [Manager] Handling: %s\n", ticket.Issue)
		fmt.Printf("  Resolution: %s\n", ticket.Description)
	} else {
		fmt.Printf("✗ [Manager] No handler available for: %s\n", ticket.Issue)
	}
}
```

## Key Components

### 1. Handler Interface
```go
type SupportHandler interface {
	SetNext(handler SupportHandler) SupportHandler
	Handle(ticket *SupportTicket)
}
```
- Defines common interface for all handlers
- `SetNext()` builds the chain (returns handler for chaining)
- `Handle()` processes or forwards the request

### 2. Base Handler
```go
type BaseSupportHandler struct {
	next SupportHandler
}

func (h *BaseSupportHandler) SetNext(handler SupportHandler) SupportHandler {
	h.next = handler
	return handler // Enables method chaining
}
```
- Provides default chaining behavior
- Eliminates duplicate code in concrete handlers
- Returns handler for fluent API

### 3. Concrete Handlers
Each handler:
- Checks if it can handle the request
- Processes it if capable
- Otherwise forwards to next handler

## Usage Example

```go
// Create handlers
level1 := &Level1Support{}
level2 := &Level2Support{}
manager := &ManagerSupport{}

// Build the chain (fluent API)
level1.SetNext(level2).SetNext(manager)

// Create and process tickets
tickets := []*SupportTicket{
	{Issue: "Password Reset", Priority: 1, Description: "..."},
	{Issue: "System Outage", Priority: 3, Description: "..."},
}

for _, ticket := range tickets {
	level1.Handle(ticket) // Always start at chain beginning
}
```

## Output Example
```
Ticket #1: Password Reset
✓ [Level 1 Support] Handling: Password Reset
  Resolution: Reset password via email link

Ticket #2: Critical System Outage
→ [Level 1 Support] Escalating: Critical System Outage
→ [Level 2 Support] Escalating: Critical System Outage
✓ [Manager] Handling: Critical System Outage
  Resolution: Escalated to infrastructure team
```

## Advantages

✅ **Reduced Coupling**: Sender doesn't need to know which handler will process the request  
✅ **Flexibility**: Easy to add/remove/reorder handlers without affecting clients  
✅ **Single Responsibility**: Each handler focuses on specific type of request  
✅ **Dynamic Chains**: Chain can be modified at runtime  
✅ **Open/Closed Principle**: New handlers can be added without modifying existing code

## Disadvantages

⚠️ **No Guarantee of Handling**: Request might reach end of chain without being processed  
⚠️ **Debugging Complexity**: Hard to trace request path through the chain  
⚠️ **Performance**: Every request traverses the chain until handled

## Best Practices

1. **Base Handler**: Use composition to avoid duplicating chain management code
2. **Clear Responsibility**: Each handler should have well-defined criteria
3. **End-of-Chain Handling**: Always handle case when no handler processes request
4. **Immutable Chain**: Consider making chain structure immutable after construction
5. **Early Exit**: Handle request as soon as possible to avoid unnecessary traversal

## Variations

### 1. Pure Chain
Each handler either processes OR forwards (never both)

### 2. Impure Chain
Handlers can process AND forward (middleware pattern)

```go
func (h *LoggingHandler) Handle(request *Request) {
	log.Printf("Processing: %v", request)
	if h.next != nil {
		h.next.Handle(request) // Always forward
	}
}
```

### 3. Branching Chain
Handlers can route to different chains based on request type

## Related Patterns

- **Command**: Chain of Responsibility can use Commands as requests
- **Composite**: Chain often used with Composite for tree structures
- **Decorator**: Similar structure, but Decorator adds behavior, CoR delegates

## Testing Strategy

```go
func TestChainOfResponsibility(t *testing.T) {
	level1 := &Level1Support{}
	level2 := &Level2Support{}
	
	level1.SetNext(level2)
	
	ticket := &SupportTicket{Priority: 2, Issue: "Test"}
	level1.Handle(ticket)
	
	// Verify level2 handled it, not level1
}
```

## Common Use Cases in Go

1. **HTTP Middleware**: 
   ```go
   r.Use(LoggingMiddleware, AuthMiddleware, RateLimitMiddleware)
   ```

2. **Error Handling**:
   ```go
   for _, handler := range errorHandlers {
       if handler.CanHandle(err) {
           return handler.Handle(err)
       }
   }
   ```

3. **Input Validation**:
   ```go
   validators := []Validator{EmailValidator, PasswordValidator, AgeValidator}
   for _, v := range validators {
       if err := v.Validate(input); err != nil {
           return err
       }
   }
   ```

## Summary

The Chain of Responsibility pattern provides a flexible way to handle requests by passing them through a chain of handlers. It's particularly useful when:
- Multiple objects can handle a request, but the handler isn't known in advance
- You want to issue a request to multiple objects without specifying the receiver explicitly
- The set of handlers should be specified dynamically

This pattern is fundamental to many frameworks and is especially common in web development (middleware), event handling, and request processing pipelines.
