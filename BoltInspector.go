package main

import (
	"bufio"
	"fmt"
	"github.com/89yoyos/Wrench"
	"os"
	"strings"
)

var w = Wrench.Wrench{}

func main() {
	prompt := true
	args := os.Args[1:]
	if len(args)!=0{
		ex, _ := FileExists(args[0])
		if ex{
			prompt = false
			err := w.OpenDB(args[0])
			if err != nil {
				fmt.Printf("Failed to open database %s",args[0])
				prompt = true
			}
		}
	}
	if prompt {
		for {
			fmt.Println()
			fmt.Print("Database to Read (or exit): ")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			DBPath := scan.Text()
			
			ex, _ := FileExists(DBPath)
			
			if DBPath == "exit" {
				fmt.Println("Exiting...")
				os.Exit(0)
				return
			} else if !ex {
				fmt.Print("[Error] The specified file does not exist.\n")
			} else {
				err := w.OpenDB(DBPath)
				if err == nil{
					break
				} else {
					fmt.Printf("[Error] Failed to open the specified file (%s)", DBPath)
				}
			}
		}
	}
	
	for {
		fmt.Printf("[%s] (%s) $>",w.FileName(),w.CurrentBucketString())
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		cmd := strings.SplitN(scan.Text()," ",2)
		
		if cmd[0]=="exit" {
			fmt.Println("Exiting...")
			break
		} else if cmd[0]=="help" {
			help(cmd)
		} else if cmd[0]=="list" {
			list(cmd)
		} else if cmd[0]=="rlist"{
			rlist(cmd)
		} else if cmd[0]=="cd"{
			cd(cmd)
		} else if cmd[0]=="print"{
			prnt(cmd)
		} else if cmd[0]=="write"{
			write(cmd)
		} else if cmd[0]=="bucket"{
			bucket(cmd)
		} else if cmd[0]=="delete"{
			dlt(cmd)
		} else if cmd[0]=="empty"{
			emptyBucket(cmd)
		} else if cmd[0]=="copy"{
			copy(cmd,false)
		} else if cmd[0]=="move"{
			copy(cmd,true)
		} else {
			fmt.Println("Unrecognized command. Type \"help\" to see commands")
		}
	}
	
}
