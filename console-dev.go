// console-dev.go
package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	// Change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Hello World!")
}
