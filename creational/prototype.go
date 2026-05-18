package creational

import "fmt"

// Prototype Pattern - Creates new objects by cloning existing ones

// Document interface defines the Clone method
type Document interface {
	Clone() Document
	Display()
	SetTitle(title string)
	SetContent(content string)
}

// BaseDocument contains common document fields
type BaseDocument struct {
	Title    string
	Content  string
	Author   string
	Template string
}

// ReportDocument is a concrete prototype for reports
type ReportDocument struct {
	BaseDocument
	ReportType string
	Charts     []string
}

// Clone creates a deep copy of the report document
func (d *ReportDocument) Clone() Document {
	// Create a new instance with copied values
	charts := make([]string, len(d.Charts))
	copy(charts, d.Charts)

	return &ReportDocument{
		BaseDocument: BaseDocument{
			Title:    d.Title,
			Content:  d.Content,
			Author:   d.Author,
			Template: d.Template,
		},
		ReportType: d.ReportType,
		Charts:     charts,
	}
}

func (d *ReportDocument) Display() {
	fmt.Println("    ┌─────────────────────────────────────────┐")
	fmt.Println("    │        REPORT DOCUMENT                  │")
	fmt.Println("    ├─────────────────────────────────────────┤")
	fmt.Printf("    │ Title:    %-30s│\n", d.Title)
	fmt.Printf("    │ Author:   %-30s│\n", d.Author)
	fmt.Printf("    │ Type:     %-30s│\n", d.ReportType)
	fmt.Printf("    │ Template: %-30s│\n", d.Template)
	fmt.Printf("    │ Content:  %-30s│\n", d.Content)
	if len(d.Charts) > 0 {
		fmt.Printf("    │ Charts:   %-30s│\n", fmt.Sprintf("%v", d.Charts))
	}
	fmt.Println("    └─────────────────────────────────────────┘")
}

func (d *ReportDocument) SetTitle(title string) {
	d.Title = title
}

func (d *ReportDocument) SetContent(content string) {
	d.Content = content
}

// InvoiceDocument is a concrete prototype for invoices
type InvoiceDocument struct {
	BaseDocument
	InvoiceNumber string
	TotalAmount   float64
	Currency      string
}

// Clone creates a deep copy of the invoice document
func (d *InvoiceDocument) Clone() Document {
	return &InvoiceDocument{
		BaseDocument: BaseDocument{
			Title:    d.Title,
			Content:  d.Content,
			Author:   d.Author,
			Template: d.Template,
		},
		InvoiceNumber: d.InvoiceNumber,
		TotalAmount:   d.TotalAmount,
		Currency:      d.Currency,
	}
}

func (d *InvoiceDocument) Display() {
	fmt.Println("    ┌─────────────────────────────────────────┐")
	fmt.Println("    │        INVOICE DOCUMENT                 │")
	fmt.Println("    ├─────────────────────────────────────────┤")
	fmt.Printf("    │ Title:    %-30s│\n", d.Title)
	fmt.Printf("    │ Author:   %-30s│\n", d.Author)
	fmt.Printf("    │ Invoice#: %-30s│\n", d.InvoiceNumber)
	fmt.Printf("    │ Template: %-30s│\n", d.Template)
	fmt.Printf("    │ Amount:   %-30s│\n", fmt.Sprintf("%.2f %s", d.TotalAmount, d.Currency))
	fmt.Printf("    │ Content:  %-30s│\n", d.Content)
	fmt.Println("    └─────────────────────────────────────────┘")
}

func (d *InvoiceDocument) SetTitle(title string) {
	d.Title = title
}

func (d *InvoiceDocument) SetContent(content string) {
	d.Content = content
}

// DocumentRegistry manages prototype instances
type DocumentRegistry struct {
	prototypes map[string]Document
}

// NewDocumentRegistry creates a new registry
func NewDocumentRegistry() *DocumentRegistry {
	return &DocumentRegistry{
		prototypes: make(map[string]Document),
	}
}

// Register adds a prototype to the registry
func (r *DocumentRegistry) Register(key string, prototype Document) {
	r.prototypes[key] = prototype
}

// Create clones a prototype from the registry
func (r *DocumentRegistry) Create(key string) Document {
	if prototype, ok := r.prototypes[key]; ok {
		return prototype.Clone()
	}
	return nil
}

// MyPrototype demonstrates the prototype pattern
func MyPrototype() {
	// Create prototype instances
	reportPrototype := &ReportDocument{
		BaseDocument: BaseDocument{
			Author:   "System",
			Template: "Standard Report Template",
		},
		ReportType: "Quarterly",
		Charts:     []string{"Sales Chart", "Revenue Chart"},
	}

	invoicePrototype := &InvoiceDocument{
		BaseDocument: BaseDocument{
			Author:   "Accounting Dept",
			Template: "Standard Invoice Template",
		},
		Currency: "USD",
	}

	// Create and populate registry
	registry := NewDocumentRegistry()
	registry.Register("report", reportPrototype)
	registry.Register("invoice", invoicePrototype)

	// Clone and customize documents
	fmt.Println("    Creating Q1 Report from prototype:")
	q1Report := registry.Create("report").(*ReportDocument)
	q1Report.SetTitle("Q1 2026 Report")
	q1Report.SetContent("Q1 financial analysis")
	q1Report.Display()

	fmt.Println()
	fmt.Println("    Creating Q2 Report from prototype:")
	q2Report := registry.Create("report").(*ReportDocument)
	q2Report.SetTitle("Q2 2026 Report")
	q2Report.SetContent("Q2 financial analysis")
	q2Report.Charts = []string{"Growth Chart", "Profit Chart"}
	q2Report.Display()

	fmt.Println()
	fmt.Println("    Creating Invoice from prototype:")
	invoice1 := registry.Create("invoice").(*InvoiceDocument)
	invoice1.SetTitle("Invoice - Client ABC")
	invoice1.InvoiceNumber = "INV-2026-001"
	invoice1.TotalAmount = 15000.00
	invoice1.SetContent("Professional services rendered")
	invoice1.Display()

	fmt.Println()
	fmt.Println("    Creating Another Invoice from prototype:")
	invoice2 := registry.Create("invoice").(*InvoiceDocument)
	invoice2.SetTitle("Invoice - Client XYZ")
	invoice2.InvoiceNumber = "INV-2026-002"
	invoice2.TotalAmount = 25000.00
	invoice2.SetContent("Software development project")
	invoice2.Display()
}
