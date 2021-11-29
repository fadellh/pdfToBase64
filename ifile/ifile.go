package ifile

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	inputDir = "./asset/"
	dir      = "./output/"
)

func Read() (*[]byte, string) {

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	firstFile := files[0]

	input, err := os.Open(inputDir + firstFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)
	content, _ := ioutil.ReadAll(reader)

	return &content, firstFile.Name()
}

func WriteBase64(content *[]byte, filename string) error {

	fmt.Println("write base 64")
	encoded := base64.StdEncoding.EncodeToString(*content)
	// fileName := time.Now().Date()
	filePath := dir + "_" + filename + ".txt"

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(encoded)

	if err != nil {
		return err
	}
	return nil
}

func Delete(filename string) {

	err := os.Remove(inputDir + filename)
	if err != nil {
		log.Fatal(err)
	}
}
