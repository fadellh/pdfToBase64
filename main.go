package main

import (
	"fmt"

	"convert-base64/ifile"
)

func main() {

	file, name := ifile.Read()
	err := ifile.WriteBase64(file, name)

	if err != nil {
		fmt.Println(err)
	}

	ifile.Delete(name)

}
