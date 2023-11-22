package main

// Processor is a high level module that used a Writer to process the data
type Processor struct {
	Writer Writer
}

func (p Processor) ProcessAndWrite(data []byte) error {
	return p.Writer.Write(data)
}
