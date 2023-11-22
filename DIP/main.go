package main

func main() {

	fw := FileWriter{FileName: "data.txt"}
	processor := Processor{Writer: fw}

	_ = processor.ProcessAndWrite([]byte("This data to be written to a file"))

	nw := NetworkWriter{Endpoint: "http:localhost:8080"}
	processor.Writer = nw
	_ = processor.ProcessAndWrite([]byte("This data to be written to a network"))
}
