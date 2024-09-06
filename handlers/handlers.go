package handlers

import (
	"fmt"
	"io"
	"os"
)

func getCharacters(filePath string) {

	file, err:=os.Open(filePath)

	if err!=nil{
		panic(err)
	}

	defer file.Close()

	total := 0
	for {
		data := make([]byte, 1024)

		len, err:=file.Read(data)

		if err==io.EOF{
			break
		}

		if err!=nil{
			panic(err)
		}

		total+=len
	}

	fmt.Printf("Total no of bytes in %s: %d\n", filePath, total)
}

func getLines(filePath string) {

	file, err:=os.Open(filePath)

	if err!=nil{
		panic(err)
	}

	defer file.Close()

	total := 0
	for {
		data := make([]byte, 1)

		len, err:=file.Read(data)

		if err==io.EOF{
			break
		}

		if err!=nil{
			panic(err)
		}
		if string(data)=="\n"{
			total+=len
		}
		
	}

	fmt.Printf("Total no of lines in %s: %d\n", filePath, max(1, total))
}