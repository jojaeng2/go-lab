package main

import (
	"flag"
	"fmt"
)

func main() {

	var uname string
	var pass string

	flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
	flag.StringVar(&pass, "p", "password", "Specify pass. Default is password")

	flag.Parse()

	if uname == "root" && pass == "password" {
		fmt.Println("Logging in")
	} else {
		fmt.Print("Invalid!!")
	}
}
