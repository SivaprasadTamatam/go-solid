package main

import "fmt"

// ISPPrinter implements BasicPrinter and Scanner interfaces
type ISPPrinter struct{}

func (p ISPPrinter) PrintDocument(doc string) {
	fmt.Println("Printing: ", doc)
}

func (p ISPPrinter) ScanDocument() {
	fmt.Println("Scanning a document")
}
