package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	gofor7 "github.com/DRK-Blutspende-BaWueHe/goforseven/src/hl7model"
)

/*
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	path := "./hl7files"
	pathWrite := path + "/WriteHL7"
	if _, err := os.Stat(pathWrite); os.IsNotExist(err) {
		os.Mkdir(pathWrite, os.ModePerm)
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		log.Println(file.Name())
		if !file.IsDir() {
			filename := fmt.Sprintf("%s/%s", path, file.Name())
			hl7message, err := ReadHL7File(filename)
			if err != nil {
				log.Println(err)
			} else {
				filename = fmt.Sprintf("%s/%s", pathWrite, createFilename(file.Name()))
				err = WriteHL7File(hl7message, filename)
			}
		}
	}
}*/

func ReadHL7File(filename string) (gofor7.HL7Message, error) {

	message, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	delimeter := gofor7.Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := gofor7.GetMessagesFromString(string(message), delimeter)
	if err != nil {
		log.Println(err)
	}
	hl7message, err := gofor7.GetHL7MessageFromMessages(messages)
	log.Println(hl7message.Header.Timestamp)

	return hl7message, err
}

func WriteHL7File(hl7message gofor7.HL7Message, filename string) error {

	delimeter := gofor7.Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	hl7messages, err := gofor7.GetStringFromHL7Messages(hl7message, delimeter)
	if err != nil {
		log.Println(err)
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	file.WriteString(hl7messages)
	file.Close()

	return nil
}

func createFilename(sourcefilename string) string {
	filename := ""
	sourcefilenameparts := strings.Split(sourcefilename, ".")
	sourcefilenameparts = append(sourcefilenameparts, sourcefilenameparts[len(sourcefilenameparts)-1])
	sourcefilenameparts[len(sourcefilenameparts)-2] = time.Now().Format("20060102_150405")
	filename = strings.Join(sourcefilenameparts, ".")
	if strings.HasSuffix(filename, ".") {
		filename = filename[:len(filename)-1]
	}
	return filename
}
