package main

import (
	"bufio"
	"fmt"
	"github.com/89yoyos/Wrench"
	"os"
)

func emptyBucket(cmd []string) {
	
	args,r := parseArguments(cmd,0)
	
	ww := Wrench.Wrench{}
	ww.OpenDB(w.File())
	ww.GoTo(w.CurrentBucket())
	
	if r==0 {
		if !ww.GoTo(w.StringToPath(args[0])) {
			fmt.Println("[Error] Specified bucket doesn't exist")
			return
		}
	} else if r==3 {
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}
	
	if !ww.Exists(){
		fmt.Println("[Error] The specified bucket does not exist")
		return
	}
	
	if ww.IsRoot() {
		fmt.Println("[Error] Cannot use Empty on root directory")
		return
	}
	
	bc,vc := ww.CountBoth()
	
	fmt.Printf("Are you sure you want to delete all %d %s and the %d %s in %s?\n",vc,valPlural(vc),bc,bcktPlural(bc),ww.CurrentBucketString())
	
	for{
		fmt.Println("Type 'yes' to continue or 'no' to cancel")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		if scan.Text() == "yes"{
			ww.Empty()
			fmt.Println("The specified bucket was emptied")
			break
		} else if scan.Text() == "no" {
			fmt.Println("The database was not changed")
			break
		}
		fmt.Print("[Error] Unknown input. ")
	}
}

func plurality(singular string, plural string, count int) string{
	if count==1{
		return singular
	}
	return plural
}

func valPlural(count int) string{
	return plurality("value","values",count)
}

func bcktPlural(count int) string{
	return plurality("bucket","buckets",count)
}