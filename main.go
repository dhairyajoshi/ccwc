package main

import (
	"ccwc/handlers"
	"ccwc/parser"
)

func main(){
	filename, args := parser.GetFileWithArgs()

	if len(args) == 0{
		for arg := range handlers.Handlerfactory{
			handler := handlers.Handlerfactory[arg]
			handler(filename)
		}
	}

	for _, arg := range args {
		handler := handlers.Handlerfactory[arg]
		handler(filename)
	}
	
}