package hl7model

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func GetStringFromHL7Messages(hl7message HL7Message, delimeter Delimeter) (string, error) {

	delimeterwork = delimeter
	messagestringhl7message := ""
	messagestring, err := getStringFromHL7Message(&hl7message.Header)
	if err == nil {
		messagestringhl7message = messagestring
		for _, hl7record := range hl7message.Records {
			refit := reflect.ValueOf(hl7record)
			structname := strings.Split(fmt.Sprintf("%v", refit.Type()), ".")[1]
			messagestring = ""
			switch structname {
			case "PID":
				hla7recordwrite, ok := hl7record.(PID)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "NTE":
				hla7recordwrite, ok := hl7record.(NTE)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "ORC":
				hla7recordwrite, ok := hl7record.(ORC)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "OBX":
				hla7recordwrite, ok := hl7record.(OBX)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "MSA":
				hla7recordwrite, ok := hl7record.(MSA)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "ERR":
				hla7recordwrite, ok := hl7record.(ERR)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "OBR":
				hla7recordwrite, ok := hl7record.(OBR)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "TCD":
				hla7recordwrite, ok := hl7record.(TCD)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "SID":
				hla7recordwrite, ok := hl7record.(SID)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "ZRI":
				hla7recordwrite, ok := hl7record.(ZRI)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "SAC":
				hla7recordwrite, ok := hl7record.(SAC)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "QPD":
				hla7recordwrite, ok := hl7record.(QPD)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "PCD":
				hla7recordwrite, ok := hl7record.(PCD)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "ZPE":
				hla7recordwrite, ok := hl7record.(PCD)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}
			case "PV1":
				hla7recordwrite, ok := hl7record.(PCD)
				if ok {
					messagestring, err = getStringFromHL7Message(&hla7recordwrite)
				}

			}
			if err == nil && len(messagestring) > 1 {
				if len(messagestringhl7message) == 0 {
					messagestringhl7message = messagestring
				} else {
					messagestringhl7message = messagestringhl7message + delimeter.CrLf + messagestring
				}
			}

		}

	}
	log.Println(messagestring)
	return messagestringhl7message, nil
}

func getStringFromHL7Message(it interface{}) (string, error) {

	messagestring := ""
	refit := reflect.ValueOf(it)
	values := []string{}

	segmentidentifier := fmt.Sprintf("%v", refit.Type())
	segmentidentifiers := strings.Split(segmentidentifier, ".")
	segmentidentifier = segmentidentifiers[len(segmentidentifiers)-1]
	values = append(values, segmentidentifier)
	if refit.Kind() == reflect.Ptr {
		refit = refit.Elem()
		refitt := refit.Type()
		for i := 0; i < refit.NumField(); i++ {
			fld := refitt.Field(i)
			tagvalue := fld.Tag.Get("hl7")
			if tagvalue != "" {
				value := fmt.Sprintf("%v", reflect.ValueOf(refit.Field(i)))
				values = addValueToValues(values, value, tagvalue)
			}

		}
	}
	messagestring = strings.Join(values, delimeterwork.Parentlevel)
	for strings.HasSuffix(messagestring, delimeterwork.Parentlevel) || strings.HasSuffix(messagestring, delimeterwork.Childlevels) {
		messagestring = messagestring[:len(messagestring)-1]
	}
	log.Println(messagestring)
	if messagestring == segmentidentifier {
		messagestring = ""
	}
	return messagestring, nil
}

func addValueToValues(values []string, value string, tagvalue string) []string {
	indices := strings.Split(tagvalue, ".")
	if len(indices) > 1 {
		parentindex, _ := strconv.Atoi(indices[1])
		for i := len(values); i < parentindex+1; i++ {
			values = append(values, "")
		}
		if len(indices) == 2 {
			values[parentindex] = value
			return values
		}
		existvalue := values[parentindex]
		childvalues := strings.Split(existvalue, delimeterwork.Childlevels)
		childindex, _ := strconv.Atoi(indices[2])
		if len(childvalues) < childindex+1 {
			for i := len(childvalues); i < childindex+1; i++ {
				childvalues = append(childvalues, "")
			}
		}
		childvalues[childindex] = value
		childvaluestring := strings.Join(childvalues, delimeterwork.Childlevels)
		values[parentindex] = childvaluestring
	}
	return values
}
