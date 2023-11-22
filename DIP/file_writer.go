package main

import "os"

// FileWriter implements the Writer interface for writing data to a file
type FileWriter struct {
	FileName string
}

func (fw FileWriter) Write(data []byte) error {
	file, err := os.Create(fw.FileName)

	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.Write(data)

	return err
}
