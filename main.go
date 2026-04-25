package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"github.com/alecthomas/kong"
)

var CLI struct {
	Init struct {
		NoGit bool `help:"flag to initiate even if no git file is found"`	
	}`cmd:"" initialize the app within the folder`

}
func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "init": runInit()

	default:panic(ctx.Command())
	}
}


func runInit() {
	if !CLI.Init.NoGit {	
		if _, err := os.Stat(".git"); errors.Is(err, os.ErrNotExist) {
			log.Fatal("Error: .git wasn't found, use --no-git to by pass")
			return
		}
	}

	if err := os.MkdirAll(".rag", 0755); err != nil {
		log.Fatalf("Error: could not create directory .rag: %v", err)
	}

	_,err :=NewDatabase(".rag/embedded")
	if err != nil {
		log.Fatalf("Error: couldn't create database, %v", err)	
	}
	fmt.Println("succesfully, created jitrag environment")
}
