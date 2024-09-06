package parser

import "os"

func GetFileWithArgs() (filename string, args []string){
	allArgs := os.Args
	
	for i, arg := range allArgs{
		if i==0{
			continue
		}
		
		if len(arg) <=3 && (arg[:len(arg)-1]=="-" || arg[:len(arg)-1]=="--"){
			args = append(args, string(arg[len(arg)-1]))
		} else {
			filename = arg
		}
	}

	return
}