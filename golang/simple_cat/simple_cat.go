package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type StatusErr struct {
	Message string
	err     error
}

func (statusErr StatusErr) Error() string {
	return statusErr.Message
}

func (statusErr StatusErr) Unwrap() error {
	return statusErr.err
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, nil, StatusErr{
			fmt.Sprintf("cannot open the file."),
			err,
		}
	}

	return file, func() {
		file.Close()
	}, err
}

func readFile() {
	if len(os.Args) < 2 {
		log.Fatalln("No File Specified")
	}

	file, closer, err := getFile(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	defer closer()

	readByte(file)
}

func readByte(file *os.File) {
	data := make([]byte, 2048)

	for {
		count, err := file.Read(data)
		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}

			break
		}
	}
}

func main() {
	readFile()
}
