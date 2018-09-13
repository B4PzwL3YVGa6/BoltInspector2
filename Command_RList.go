package main

import (
	"github.com/89yoyos/Wrench"
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

func rlist(cmd []string) {
	if !w.Exists(){
		fmt.Println("[Error] Bucket does not exist. Returning to root...")
		w.Reset()
		return
	}
	verbose := false
	maxRecurse := 3
	
	args,r := parseArguments(cmd,0)
	
	if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	} else if r==0 {
		for i:=0;i<len(args);i++{
			if args[i]=="v" || args[i]=="-v" {
				verbose=true
			} else if strings.HasPrefix(args[i],"d=")||strings.HasPrefix(args[i],"depth="){
				re := regexp.MustCompile("[^\\d]")
				tmp := re.ReplaceAllString(args[i],"")
				in,er:=strconv.Atoi(tmp)
				if er!= nil{
					fmt.Println("[Error] Invalid length passed")
					return
				}
				maxRecurse=in
			}
		}
	}
	
	recurseDump(w,maxRecurse,0,verbose)
	
}

func recurseDump(w Wrench.Wrench, rMax int, rCur int, verbose bool){
	bckts,vals := w.GetBoth()
	rpath := "."
	if w.IsRoot(){
		rpath = "~"
	}
	for i,val := range w.Path(){
		if i>=len(w.Path()){
			rpath=rpath+"/"+val
		}
	}
	for _,val := range vals{
		if !verbose{fmt.Print(multitab(rCur)+rpath+"/")}
		fmt.Println(val.ToString(verbose))
	}
	for _,val := range bckts{
		if !verbose{fmt.Print(multitab(rCur)+rpath+"/")}
		fmt.Print(val.ToString(verbose))
		if verbose && rMax<=rCur {
			fmt.Printf("- Contents Outside Recurse Depth (Depth = %d)\n\n", rMax)
		} else if rMax<=rCur {
			fmt.Println()
		} else if rMax>rCur {
			fmt.Println()
			ww := val.AsWrench()
			recurseDump(ww, rMax, rCur+1, verbose)
		}
	}
}