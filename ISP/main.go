package main

func main() {
	printer := ISPPrinter{}

	BasicPrinterClient(printer)

	MediumPrinterClient(printer, printer)
}
