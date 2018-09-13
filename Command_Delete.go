package main

import (
	"bufio"
	"fmt"
	"github.com/89yoyos/Wrench"
	"os"
)

func dlt(cmd []string){
	
	args,r := parseArguments(cmd,1)
	
	if r==2 {
		fmt.Println("[Error] You must specify a key to dlt")
		return
	} else if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}
	
	path := w.StringToPath(args[0])
	trgt := path[len(path)-1]
	
	ww := Wrench.Wrench{}
	ww.OpenDB(w.File())
	ww.GoTo(path[:len(path)-1])
	
	if !ww.Exists(){
		fmt.Println("[Error] The specified bucket doesn't exist")
		return
	}
	
	val,suc := ww.Get(trgt)
	
	if !suc {
		fmt.Printf("[Error] Value for key %s is undefined in %s\n",args[0],ww.CurrentBucketString())
		return
	}
	
	if val.IsBucket(){
		fmt.Printf("Are you sure you wish to dlt the BUCKET %s?\n",val.Key())
	} else {
		fmt.Printf("Are you sure you wish to dlt the KEY/VALUE %s?\n",val.Key())
	}
	
	for{
		fmt.Println("Type 'yes' to continue or 'no' to cancel")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		if scan.Text() == "yes"{
			ww.Delete(trgt)
			fmt.Printf("%s was deleted\n",val.Key())
			break
		} else if scan.Text() == "no" {
			fmt.Println("The database was not changed")
			break
		}
		fmt.Print("[Error] Unknown input. ")
	}
	
}
