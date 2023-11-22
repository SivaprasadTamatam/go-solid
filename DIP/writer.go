package main

// Writer defines the contract for writing a data
type Writer interface {
	Write(data []byte) error
}
