package main

import (
	"fmt"
	"github.com/89yoyos/Wrench"
)

func copy(cmd []string, del bool){
	
	args,r := parseArguments(cmd,2)
	
	if r==2{
		fmt.Println("[Error] You must enter both a target and destination key")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}
	
	path1 := w.StringToPath(args[0])
	path2 := w.StringToPath(args[1])
	
	k1 := path1[len(path1)-1]
	k2 := path2[len(path2)-1]
	
	
	if len(path1)<2 || len(path2)<2{
		fmt.Println("[Error] Invalid path(s) entered [length error]")
		return
	}
	
	ww1 := Wrench.Wrench{}
	ww1.OpenDB(w.File())
	ww1.GoTo(path1[:len(path1)-1])
	
	ww2 := Wrench.Wrench{}
	ww2.OpenDB(w.File())
	ww2.GoTo(path2[:len(path2)-1])
	
	if !ww1.Exists() {
		fmt.Printf("[Error] Source bucket (%s) doesn't exist\n",path1)
		return
	}
	
	if !ww2.Exists() {
		fmt.Printf("[Error] Destination bucket (%s) doesn't exist\n",path2)
	}
	
	_,e2 := ww2.Get(k2)
	
	if e2{
		fmt.Println("[Error] Destination key already exists")
		return
	}
	
	v1,e1 := ww1.Get(k1)
	
	if !e1{
		fmt.Println("[Error] Source key doesn't exist")
		return
	}
	
	if v1.IsBucket(){
		fmt.Println("[Error] Source is a bucket. Copying buckets is not yet supported")
		return
	} else {
		ww2.Insert(k2,v1.V())
	}
	
	if del{
		ww1.Delete(v1.Key())
	}
	
}