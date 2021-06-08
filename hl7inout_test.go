package hl7inout

import (
	"fmt"
	"testing"
	//gofor7 "github.com/DRK-Blutspende-BaWueHe/goforseven/hl7inout"
)

func TestMSH(t *testing.T) {
	MSHstringSource := `MSH|^~\&|CIT|LAB|HL7_HOST|HL7_OFFICE|20120119180929||ORU^R01|248914|P|2.3|||NE|NE||8859/1`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(MSHstringSource, delimeter)
	if err != nil {
		t.Errorf("TestMSH returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	MSHstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if MSHstringSource != MSHstringCreated {
		t.Errorf(`TestMSH wrong MSH want: %v get: %v`, MSHstringSource, MSHstringCreated)
	}
}

func TestPID(t *testing.T) {
	PIDstringSource := `PID|1||400411|400412|Boop^Elizabeth||19301213|F`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(PIDstringSource, delimeter)
	if err != nil {
		t.Errorf("TestPID returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	PIDstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if PIDstringSource != PIDstringCreated {
		t.Errorf(`TestPID wrong PID want: %v get: %v`, PIDstringSource, PIDstringCreated)
	}

}

func TestORC(t *testing.T) {
	ORCstringSource := `ORC|RE|180027|||||^^^^^R||20120119171240|||Thanh`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(ORCstringSource, delimeter)
	if err != nil {
		t.Errorf("TestORC returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	ORCstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if ORCstringSource != ORCstringCreated {
		t.Errorf(`TestORC wrong ORC want: %v get: %v`, ORCstringSource, ORCstringCreated)
	}

}

func TestOBR(t *testing.T) {
	OBRstringSource := `OBR|1|180027||ALL||20120119171240|||||||^^^||S1^^^^^^P`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(OBRstringSource, delimeter)
	if err != nil {
		t.Errorf("TestOBR returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	OBRstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if OBRstringSource != OBRstringCreated {
		t.Errorf(`TestOBR wrong OBR want: %v get: %v`, OBRstringSource, OBRstringCreated)
	}

}

func TestOBX(t *testing.T) {
	OBXstringSource := `OBX|1||21||111.0|mmol/L|^^|H|||P||N|20120119171543|^^^COBAS8K.200|System`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(OBXstringSource, delimeter)
	if err != nil {
		t.Errorf("TestOBX returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	OBXstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if OBXstringSource != OBXstringCreated {
		t.Errorf(`TestOBX wrong OBX want: %v get: %v`, OBXstringSource, OBXstringCreated)
	}

}

func TestTCD(t *testing.T) {
	TCDstringSource := `TCD|1|21|Inc`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(TCDstringSource, delimeter)
	if err != nil {
		t.Errorf("TestTCD returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	TCDstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if TCDstringSource != TCDstringCreated {
		t.Errorf(`TestTCD wrong TCD want: %v get: %v`, TCDstringSource, TCDstringCreated)
	}

}

func TestNTE(t *testing.T) {
	NTEstringSource := `NTE||L|1st·comment·on·patient·/·sample·20020604101`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(NTEstringSource, delimeter)
	if err != nil {
		t.Errorf("TestNTE returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	NTEstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if NTEstringSource != NTEstringCreated {
		t.Errorf(`TestNTE wrong NTE want: %v get: %v`, NTEstringSource, NTEstringCreated)
	}

}

func TestPV1(t *testing.T) {
	PV1stringSource := `PV1|^^^^^^^^^||||||||||||||||||case0815`
	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	messages, err := GetMessagesFromString(PV1stringSource, delimeter)
	if err != nil {
		t.Errorf("TestPV1 returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	PV1stringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if PV1stringSource != PV1stringCreated {
		t.Errorf(`TestPV1 wrong PV1 want: %v get: %v`, PV1stringSource, PV1stringCreated)
	}

}

func TestComplett(t *testing.T) {
	MSHstringSource := `MSH|^~\&|CIT|LAB|HL7_HOST|HL7_OFFICE|20120119180929||ORU^R01|248914|P|2.3|||NE|NE||8859/1`
	PIDstringSource := `PID|1||400411|400412|Boop^Elizabeth||19301213|F`
	PV1stringSource := `PV1|^^^^^^^^^||||||||||||||||||case0815`
	NTEstringSource := `NTE||L|1st·comment·on·patient·/·sample·20020604101`
	ORCstringSource := `ORC|RE|180027|||||^^^^^R||20120119171240|||Thanh`
	OBRstringSource := `OBR|1|180027||ALL||20120119171240|||||||^^^||S1^^^^^^P`
	OBXstringSource := `OBX|1||21||111.0|mmol/L|^^|H|||P||N|20120119171543|^^^COBAS8K.200|System`
	TCDstringSource := `TCD|1|21|Inc`

	delimeter := Delimeter{}
	delimeter.Parentlevel = "|"
	delimeter.Childlevels = "^"
	delimeter.CrLf = "\r"
	ComplettstringSource := fmt.Sprintf("%s%s%s", MSHstringSource, delimeter.CrLf, PIDstringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, PV1stringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, NTEstringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, ORCstringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, OBRstringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, OBXstringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, TCDstringSource)
	ComplettstringSource = fmt.Sprintf("%s%s%s", ComplettstringSource, delimeter.CrLf, ORCstringSource)
	messages, err := GetMessagesFromString(ComplettstringSource, delimeter)
	if err != nil {
		t.Errorf("Test returned err: %v", err)
	}
	hl7messageSource, err := GetHL7MessageFromMessages(messages)
	ComplettstringCreated, err := GetStringFromHL7Messages(hl7messageSource, delimeter)
	if ComplettstringSource != ComplettstringCreated {
		t.Errorf(`Test wrong want: %v get: %v`, ComplettstringSource, ComplettstringCreated)
	}

}
