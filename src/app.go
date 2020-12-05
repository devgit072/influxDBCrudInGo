package main

import "log"

func main() {
	if err := writeSomeData(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := readData(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
