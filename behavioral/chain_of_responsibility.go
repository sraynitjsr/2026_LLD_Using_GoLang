package behavioral

import "fmt"

// Chain of Responsibility Pattern - Passes requests along a chain of handlers

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
		fmt.Printf("    ✓ [Level 1 Support] Handling: %s\n", ticket.Issue)
		fmt.Printf("      Resolution: %s\n", ticket.Description)
		fmt.Println()
	} else if h.next != nil {
		fmt.Printf("    → [Level 1 Support] Escalating: %s (Priority %d)\n", ticket.Issue, ticket.Priority)
		h.next.Handle(ticket)
	}
}

// Level2Support handles medium complexity issues (Priority 2)
type Level2Support struct {
	BaseSupportHandler
}

func (h *Level2Support) Handle(ticket *SupportTicket) {
	if ticket.Priority == 2 {
		fmt.Printf("    ✓ [Level 2 Support] Handling: %s\n", ticket.Issue)
		fmt.Printf("      Resolution: %s\n", ticket.Description)
		fmt.Println()
	} else if h.next != nil {
		fmt.Printf("    → [Level 2 Support] Escalating: %s (Priority %d)\n", ticket.Issue, ticket.Priority)
		h.next.Handle(ticket)
	}
}

// ManagerSupport handles high priority issues (Priority 3)
type ManagerSupport struct {
	BaseSupportHandler
}

func (h *ManagerSupport) Handle(ticket *SupportTicket) {
	if ticket.Priority == 3 {
		fmt.Printf("    ✓ [Manager] Handling: %s\n", ticket.Issue)
		fmt.Printf("      Resolution: %s\n", ticket.Description)
		fmt.Println()
	} else if h.next != nil {
		fmt.Printf("    → [Manager] Cannot handle: %s\n", ticket.Issue)
		h.next.Handle(ticket)
	} else {
		fmt.Printf("    ✗ [Manager] No handler available for: %s\n", ticket.Issue)
		fmt.Println()
	}
}

// MyChainOfResponsibility demonstrates the chain of responsibility pattern
func MyChainOfResponsibility() {
	fmt.Println("    Building Support Chain: Level1 → Level2 → Manager")
	fmt.Println()

	// Create handlers
	level1 := &Level1Support{}
	level2 := &Level2Support{}
	manager := &ManagerSupport{}

	// Build the chain
	level1.SetNext(level2).SetNext(manager)

	// Create various support tickets
	tickets := []*SupportTicket{
		{
			Issue:       "Password Reset",
			Priority:    1,
			Description: "Reset password via email link",
		},
		{
			Issue:       "Software Installation Issue",
			Priority:    2,
			Description: "Provided installation guide and troubleshooting steps",
		},
		{
			Issue:       "Critical System Outage",
			Priority:    3,
			Description: "Escalated to infrastructure team, monitoring resolution",
		},
		{
			Issue:       "Email Configuration",
			Priority:    1,
			Description: "Updated SMTP settings in user profile",
		},
	}

	// Process tickets through the chain
	for i, ticket := range tickets {
		fmt.Printf("    Ticket #%d: %s\n", i+1, ticket.Issue)
		level1.Handle(ticket)
	}

	fmt.Println("    Chain of Responsibility: Pass request along handler chain")
}
