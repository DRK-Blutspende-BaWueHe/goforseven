package hl7model

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func GetMessagesFromString(hl7string string, delimeter Delimeter) ([]string, error) {

	delimeterwork = delimeter

	messagesreturn := []string{}
	hl7string = strings.ReplaceAll(hl7string, "\n", "\r")
	messages := strings.Split(hl7string, "\r")

	for _, message := range messages {
		if len(message) > 0 {
			message = strings.ReplaceAll(message, "\t", "")
			message = strings.TrimLeft(message, " ")
			messagesreturn = append(messagesreturn, message)
		}
	}
	return messagesreturn, nil
}

func GetHL7MessageFromMessages(messages []string) (HL7Message, error) {

	hl7message := HL7Message{}

	for _, message := range messages {
		switch strings.Split(message, delimeterwork.Parentlevel)[0] {
		case "MSH":
			msh := MSH{}
			err := unmarshal(message, &msh)
			if err != nil {
				log.Println(err)
			}
			hl7message.Header = msh
		case "PID":
			hl7 := PID{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "NTE":
			hl7 := NTE{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "ORC":
			hl7 := ORC{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "OBX":
			hl7 := OBX{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "MSA":
			hl7 := MSA{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "ERR":
			hl7 := ERR{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "OBR":
			hl7 := OBR{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "TCD":
			hl7 := TCD{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "EQU":
			hl7 := EQU{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "SID":
			hl7 := SID{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "ZRI":
			hl7 := ZRI{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "SAC":
			hl7 := SAC{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "QPD":
			hl7 := QPD{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "PCD":
			hl7 := PCD{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "ZPE":
			hl7 := ZPE{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		case "PV1":
			hl7 := PV1{}
			err := unmarshal(message, &hl7)
			if err != nil {
				log.Println(err)
			}
			hl7message.Records = append(hl7message.Records, hl7)
		}
	}

	return hl7message, nil
}

func unmarshal(message string, it interface{}) error {
	log.Println(message)
	refit := reflect.ValueOf(it) // .Elem()
	if refit.Kind() == reflect.Ptr {
		refit = refit.Elem()
		refitt := refit.Type()
		for i := 0; i < refit.NumField(); i++ {
			fld := refitt.Field(i)
			tagvalue := fld.Tag.Get("hl7")
			if tagvalue != "" {
				indices := strings.Split(tagvalue, ".")
				value := ""
				if len(indices) > 1 {
					value = message
					delimeter := delimeterwork.Parentlevel
					for j := 1; j < len(indices); j++ {
						if j > 1 {
							delimeter = delimeterwork.Childlevels
							if !strings.Contains(value, delimeterwork.Childlevels) {
								index, _ := strconv.Atoi(indices[j])
								if index > 0 {
									value = ""
								}
								break
							}
						}
						index, _ := strconv.Atoi(indices[j])
						value = getIndexValue(value, delimeter, index)
					}

				}
				switch refit.Field(i).Kind() {
				case reflect.String:
					refit.Field(i).SetString(strings.TrimSpace(value))
				case reflect.Int:
					intvalue, err := strconv.ParseInt(value, 10, 64)
					if err == nil {
						refit.Field(i).SetInt(intvalue)
					}
				case reflect.Float64:
					floatvalue, err := strconv.ParseFloat("value", 64)
					if err == nil {
						refit.Field(i).SetFloat(floatvalue)
					}
				case reflect.Bool:
					boolvalue, err := strconv.ParseBool("value")
					if err == nil {
						refit.Field(i).SetBool(boolvalue)
					}
				}
			}
		}
	}
	return nil
}

func getIndexValue(value string, delimeter string, index int) string {
	valuefields := strings.Count(value, delimeter)
	indexvalue := ""
	if valuefields >= index {
		indexvalue = strings.Split(value, delimeter)[index]
	}
	return indexvalue
}
