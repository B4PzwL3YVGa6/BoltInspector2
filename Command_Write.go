package main

import (
	"fmt"
	"encoding/binary"
	"github.com/89yoyos/Wrench"
	"strconv"
	"bytes"
)

func write(cmd []string){
	
	if len(w.Path())==1{
		fmt.Println("[Error] You cannot write values in the root directory. Use the 'bucket' command to make a value storing bucket.")
		return
	}
	
	args,r := parseArguments(cmd,2)
	
	if r==2{
		fmt.Println("[Error] You must specify a key and value to write.")
		return
	} else if r==3{
		fmt.Println("[Error] Couldn't parse arguments")
		return
	}
	
	path := w.StringToPath(args[0])
	trgt := path[len(path)-1]
	
	ww := Wrench.Wrench{}
	ww.OpenDB(w.File())
	ww.GoTo(path[:len(path)-1])
	
	if !ww.Exists() {
		fmt.Println("[Error] The specified bucket doesn't exist")
		return
	}
	
	dt:="string"
	
	if len(args)>2{
		if args[2]=="s" || args[2]=="s"{
			dt="string"
		} else if args[2]=="i" || args[2]=="-i" {
			dt="int"
		} else if args[2]=="ui" || args[2]=="-ui"{
			dt="uint"
		}
	}
	
	if dt=="string"{
		ww.Insert(trgt,[]byte(args[1]))
	} else if dt=="int" {
		i, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			fmt.Printf("[Error] Couldn't convert value string to int:\n%s\n", err)
			return
		}
		buf := new(bytes.Buffer)
		err2 := binary.Write(buf, binary.LittleEndian, i)
		if err2 != nil {
			fmt.Println("binary.Write failed:", err2)
			return
		}
		ww.Insert(trgt, buf.Bytes())
		return
	} else if dt=="uint" {
		i, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			fmt.Printf("[Error] Couldn't convert value string to int:\n%s", err)
			return
		}
		buf := new(bytes.Buffer)
		err2 := binary.Write(buf, binary.LittleEndian, i)
		if err2 != nil {
			fmt.Println("[Error] binary.Write failed:", err2)
			return
		}
		ww.Insert(trgt, buf.Bytes())
		return
	} else {
		fmt.Println("[Error] Unknown Insert Type.")
		return
	}
}