package main

// Printer defines contract for a printer

type Printer interface {
	PrintDocument(doc string)
	ScanDocument()
	FaxDocument()
}
