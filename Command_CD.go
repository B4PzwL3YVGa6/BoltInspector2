package main

import (
	"fmt"
)

func cd(cmd []string){
	
	args,r := parseArguments(cmd,1)
	
	if r==2{
		fmt.Println("[Error] You must enter a destination")
	}
	
	bp := w.StringToPath(args[0])
	
	if !w.GoTo(bp){
		fmt.Println("[Error] Invalid bucket path: " + args[0])
	}
	
}