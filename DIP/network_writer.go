package main

import "fmt"

// NetworkWriter implements the Writer interface for sending data over the network

type NetworkWriter struct {
	Endpoint string
}

func (nw NetworkWriter) Write(data []byte) error {
	fmt.Printf("Sending data  %s to %s", string(data), nw.Endpoint)
	return nil
}
