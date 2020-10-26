package main

import (
	s "strings"
	"fmt"
)

var p = fmt.Println

func main() { 

	p("Contains:  ", s.Contains("test", "es"))
	
}

