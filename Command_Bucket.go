package main

import (
	"fmt"
)

func bucket(cmd []string){
	
	args,r := parseArguments(cmd,1)
	
	if r==2{
		fmt.Println("[Error] You must specify a key for the bucket.")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments.")
		return
	}
	
	_,suc := w.Get(args[0])
	
	if suc {
		fmt.Println("[Error] Value for key \"" + args[0] + "\" is already defined in this bucket (" + w.CurrentBucketString() + ")")
		return
	}
	
	w.CreateBucket(args[0])
	
}
