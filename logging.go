package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")

	// Println + os.Exit(1)
	log.Fatalln("fatal message")

	// Println + panic()
	log.Panicln("panic message")
}
