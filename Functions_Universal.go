package main

import (
	"github.com/mattn/go-shellwords"
	"os"
)

// Check if file FileExists at a given path, and that the path is valid
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func parseArguments(cmd []string, minArgs int) (args []string, returnCode int){
	
	hasArgs := len(cmd)>1
	
	if !hasArgs && minArgs<=0{
		return []string{},1 // No arguments
	} else if !hasArgs && minArgs>0{
		return []string{},2 // Not enough arguments
	}
	
	args,err := shellwords.Parse(cmd[1])
	
	if err!=nil{
		return []string{},3 // Failed to parse
	}
	
	if len(args)<minArgs {
		return []string{},2 // Not enough arguments
	}
	
	return args,0
}

func multitab(n int) string{
	if n<0{
		return ""
	}
	s := ""
	for i:=0;i<n;i++{
		s=s+"\t"
	}
	return s
}