package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var ConfigPath string = "test.pro"

func ChangeTestConfig(address, port, dbname string) error {
	file, err := os.OpenFile(ConfigPath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open config file fail, err: ", err)
		return err
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	output := make([]byte, 0)
	for {
		line, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(line), "address") {
			newline := "address = " + address
			line = []byte(newline)
		}
		if strings.Contains(string(line), "port") {
			newline := "port = " + port
			line = []byte(newline)
		}
		if strings.Contains(string(line), "dbname") {
			newline := "dbname = " + dbname
			line = []byte(newline)
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}

	if err := writeToFile(ConfigPath, output); err != nil {
		fmt.Println("write config file err: %v", err)
		return err
	}
	return nil
}

func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func main() {
	ChangeTestConfig("192.168.33.216", "3306", "mdm")
}
