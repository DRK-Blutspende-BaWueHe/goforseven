package hl7inout

type HL7Message struct {
	Header  MSH
	Records []interface{}
}

// MSH - message header segment
// MSH|^~\&|||||20160414174306||ORU^R01|234057|P|2.3|||NE|NE||8859/1
type MSH struct {
	FieldSeparator            string `hl7:"MSH.1"` // 1
	EncodingCharacter         string `hl7:"MSH.2"` // 2
	SendingApplication        string `hl7:"MSH.3"`
	SendingFacility           string `hl7:"MSH.4"`  // 4
	ReceivingApplication      string `hl7:"MSH.5"`  // 5
	ReceivingFacility         string `hl7:"MSH.6"`  // 6
	Timestamp                 string `hl7:"MSH.7"`  // 7
	Security                  string `hl7:"MSH.8"`  // 8	not used
	MessageType               string `hl7:"MSH.9"`  // 9
	MessageControlID          string `hl7:"MSH.10"` // 10
	ProcessingID              string `hl7:"MSH.11"` // 11
	VersionID                 string `hl7:"MSH.12"` // 12
	SequenceNumber            string `hl7:"MSH.13"` // 13	not used
	ContinuationPointer       string `hl7:"MSH.14"` // 14
	AcceptAcknowledgment      string `hl7:"MSH.15"` // 15
	ApplicationAcknowledgment string `hl7:"MSH.16"` // 16
	CountryCode               string `hl7:"MSH.17"` // 17	not used
	CharacterSet              string `hl7:"MSH.18"` // 18
	Language                  string `hl7:"MSH.19"` // 19	not used
}

// MSA - message acknowledgment segment
type MSA struct {
	AcknowledgmentCode         string `hl7:"MSA.1"` // 1
	MessageControlID           string `hl7:"MSA.2"` // 2
	TextMessage                string `hl7:"MSA.3"` // 3
	ExpectedSequenceNumber     string `hl7:"MSA.4"` // 4	not used
	DelayedAcknowledgementType string `hl7:"MSA.5"` // 5	not used
	ErrorCondition             string `hl7:"MSA.6"` // 6	not used
}

// ERR - error segment
type ERR struct {
	TextMessage string `hl7:"ERR.1"` // 1
}

// NTE - notes and comments segment
// NTE||L|1st·comment·on·patient·/·sample·20020604101
type NTE struct {
	SetIDNTE        string `hl7:"NTE.1"`   // 1
	SourceofComment string `hl7:"NTE.2"`   // 2
	ResultStatus    string `hl7:"NTE.3.0"` // 3.0
	ManualStatus    string `hl7:"NTE.3.1"` // 3.1
	Comment1        string `hl7:"NTE.4"`   // 4
	Comment2        string `hl7:"NTE.5"`   // 5
}

// PID - patient identification segment
type PID struct {
	PatientID              string `hl7:"PID.1"`    // 1
	ExternalPatientID      string `hl7:"PID.2"`    // 2
	InternalPatientID      string `hl7:"PID.3"`    // 3
	AlternatePatientID     string `hl7:"PID.4"`    // 4
	Surname                string `hl7:"PID.5.0"`  // 5
	FirstName              string `hl7:"PID.5.1"`  // 5
	PatientMaidenName      string `hl7:"PID.6"`    // 6		not used
	DOB                    string `hl7:"PID.7"`    // 7
	Gender                 string `hl7:"PID.8"`    // 8
	PatientAlias           string `hl7:"PID.9"`    // 9
	Race                   string `hl7:"PID.10"`   // 10
	Street                 string `hl7:"PID.11.0"` // 11.1
	City                   string `hl7:"PID.11.2"` // 11.3
	State                  string `hl7:"PID.11.3"` // 11.4
	CountyCode             string `hl7:"PID.12"`   // 12	not used
	PhoneNumberHome        string `hl7:"PID.13"`   // 13	not used
	PhoneNumberBusiness    string `hl7:"PID.14"`   // 14	not used
	PrimaryLanguage        string `hl7:"PID.15"`   // 15	not used
	Marital                string `hl7:"PID.16"`   // 16	not used
	Religion               string `hl7:"PID.17"`   // 17	not used
	PatientAccountNumber   string `hl7:"PID.18"`   // 18	not used
	SSNNumberPatient       string `hl7:"PID.19"`   // 19
	DriveLicenseNumber     string `hl7:"PID.20"`   // 20	not used
	MotherIdentifier       string `hl7:"PID.21"`   // 21	not used
	EthnicGroup            string `hl7:"PID.22"`   // 22	not used
	BirthPlace             string `hl7:"PID.23"`   // 23	not used
	MultipleBirthIndicator string `hl7:"PID.24"`   // 24	not used
	BirthOrder             string `hl7:"PID.25"`   // 25	not used
	Citizenship            string `hl7:"PID.26"`   // 26	not used
	VeteransMilitary       string `hl7:"PID.27"`   // 27	not used
	Nationality            string `hl7:"PID.28"`   // 28	not used
	PatientDeathDate       string `hl7:"PID.29"`   // 29	not used
	PatientDeathIndicator  string `hl7:"PID.30"`   // 30	not used
}

// ORC - common order segment
type ORC struct {
	OrderControl            string `hl7:"ORC.1"`    // 1
	PlacerOrderNumber       string `hl7:"ORC.2"`    // 2
	FillerOrderNumber       string `hl7:"ORC.3"`    // 3		not used
	PlacerGroupNumber       string `hl7:"ORC.4"`    // 4
	OrderStatus             string `hl7:"ORC.5"`    // 5
	ResponseFlag            string `hl7:"ORC.6"`    // 6		not used
	Priorityflag1           string `hl7:"ORC.7.0"`  // 7.1	not used
	Priorityflag2           string `hl7:"ORC.7.1"`  // 7.2	not used
	Priorityflag3           string `hl7:"ORC.7.2"`  // 7.3	not used
	Priorityflag4           string `hl7:"ORC.7.3"`  // 7.4	not used
	Priorityflag5           string `hl7:"ORC.7.4"`  // 7.5	not used
	Priorityflag6           string `hl7:"ORC.7.5"`  // 7.6
	Parent                  string `hl7:"ORC.8"`    // 8		not used
	TransactionTimestamp    string `hl7:"ORC.9"`    // 9
	EnteredBy               string `hl7:"ORC.10"`   // 10 	not used
	VerifiedBy              string `hl7:"ORC.11"`   // 11 	not used
	PhysicianCode           string `hl7:"ORC.12.0"` // 12.1
	PhysicianSurname        string `hl7:"ORC.12.1"` // 12.2
	PhysicianFirstname      string `hl7:"ORC.12.2"` // 12.3
	EntererLocation         string `hl7:"ORC.13"`   // 13 	not used
	CallBackPhoneNumber     string `hl7:"ORC.14"`   // 14 	not used
	OrderEffectiveTimestamp string `hl7:"ORC.15"`   // 15 	not used
	OrderControlCodeReason  string `hl7:"ORC.16"`   // 16
	EnteringOrganization    string `hl7:"ORC.17"`   // 17
	EnteringDevice          string `hl7:"ORC.18"`   // 18 	not used
	ActionBy                string `hl7:"ORC.19"`   // 19 	not used
}

// OBR - observation request segment
type OBR struct {
	SetID                              string `hl7:"OBR.1"`    // 1
	PlacerOrderNumber                  string `hl7:"OBR.2"`    // 2
	FillerOrderNumber                  string `hl7:"OBR.3"`    // 3
	UniversalServiceID                 string `hl7:"OBR.4"`    // 4
	Priority                           string `hl7:"OBR.5"`    // 5
	RequestedTimestamp                 string `hl7:"OBR.6"`    // 6
	ObservationTimestamp               string `hl7:"OBR.7"`    // 7		not supported
	ObservationEndTimestamp            string `hl7:"OBR.8"`    // 8		not supported
	CollectionVolume                   string `hl7:"OBR.9"`    // 9		not supported
	CollectorIdentifier                string `hl7:"OBR.10"`   // 10	not supported
	ActionCode                         string `hl7:"OBR.11"`   // 11
	DangerCode                         string `hl7:"OBR.12"`   // 12 	not used
	Pregnancy                          string `hl7:"OBR.13.0"` // 13.1
	MenstrualCycle                     string `hl7:"OBR.13.1"` // 13.2
	Medication                         string `hl7:"OBR.13.2"` // 13.3
	DiagnosticServSectID               string `hl7:"OBR.13.3"` // 13.4
	ReceivedTimestamp                  string `hl7:"OBR.14"`   // 14 	not used
	SampleType1                        string `hl7:"OBR.15.0"` // 15.1
	SampleType2                        string `hl7:"OBR.15.1"` // 15.2		not use
	SampleType3                        string `hl7:"OBR.15.2"` // 15.3		not use
	SampleType4                        string `hl7:"OBR.15.3"` // 15.4		not use
	SampleType5                        string `hl7:"OBR.15.4"` // 15.5		not use
	SampleType6                        string `hl7:"OBR.15.5"` // 15.6		not use
	SampleRole                         string `hl7:"OBR.15.6"` // 15.7
	OrderingProviderPhysicianCode      string `hl7:"OBR.16.0"` // 16.1
	OrderingProviderPhysicianSurName   string `hl7:"OBR.16.1"` // 16.2
	OrderingProviderPhysicianFirstName string `hl7:"OBR.16.2"` // 16.3
	OrderCallbackNumber                string `hl7:"OBR.17"`   // 17	not used
	PlacerField1                       string `hl7:"OBR.18"`   // 18
	PlacerField2                       string `hl7:"OBR.19"`   // 19
	FillerField                        string `hl7:"OBR.20"`   // 20 	not used
	SampleQuality                      string `hl7:"OBR.21"`   // 21 	not used
	ResultsChangeTime                  string `hl7:"OBR.22"`   // 22
	ChargetoPractice                   string `hl7:"OBR.23"`   // 23
	DiagnosticServID                   string `hl7:"OBR.24"`   // 24
	ResultStatus                       string `hl7:"OBR.25"`   // 25 	not used
	ParentResult                       string `hl7:"OBR.26"`   // 26 	not used
	PriorityCode1                      string `hl7:"OBR.27.0"` // 27.1	not use
	PriorityCode2                      string `hl7:"OBR.27.1"` // 27.2	not use
	PriorityCode3                      string `hl7:"OBR.27.2"` // 27.3	not use
	PriorityCode4                      string `hl7:"OBR.27.3"` // 27.4	not use
	PriorityCode5                      string `hl7:"OBR.27.4"` // 27.5	not use
	PriorityCode6                      string `hl7:"OBR.27.5"` // 27.6
	ResultCopiesTo                     string `hl7:"OBR.28"`   // 28 	not used
	Parent                             string `hl7:"OBR.29"`   // 29 	not used
	TransportationMode                 string `hl7:"OBR.30"`   // 30 	not used
	ReasonforStudy                     string `hl7:"OBR.31"`   // 31 	not used
	PrincipalResult                    string `hl7:"OBR.32"`   // 32 	not used
	AssistantResult                    string `hl7:"OBR.33"`   // 33 	not used
	Technician                         string `hl7:"OBR.34"`   // 34 	not used
	Transcriptionist                   string `hl7:"OBR.35"`   // 35 	not used
	ScheduledTimestamp                 string `hl7:"OBR.36"`   // 36 	not used
	NumberSampleContainers             string `hl7:"OBR.37"`   // 37 	not used
	TransportLogistics                 string `hl7:"OBR.38"`   // 38 	not used
	CollectorComment                   string `hl7:"OBR.39"`   // 39 	not used
	TransportArrangement               string `hl7:"OBR.40"`   // 40 	not used
	TransportArranged                  string `hl7:"OBR.41"`   // 41 	not used
	EscortRequired                     string `hl7:"OBR.42"`   // 42 	not used
	PlannedPatientTransport            string `hl7:"OBR.43"`   // 43 	not used
}

// TCD - test code details segment
type TCD struct {
	SetID                      string `hl7:"TCD.1"`   // 1
	UniversalServiceIdentifier string `hl7:"TCD.2"`   // 2
	InstrumentDilution         string `hl7:"TCD.3"`   // 3
	RerunDilutionFactor        string `hl7:"TCD.4"`   // 4		not used
	PreDilutionFactor          string `hl7:"TCD.5"`   // 5 	not used
	EndogenousContent          string `hl7:"TCD.6"`   // 6 	not used
	AutomaticRerunAllowed      string `hl7:"TCD.7"`   // 7 	not used
	ReflexAllowed              string `hl7:"TCD.8"`   // 8
	PipettingVolume            string `hl7:"TCD.9.0"` // 9.1
	PipettingVolumeUnit        string `hl7:"TCD.9.1"` // 9.2
	PoolingSize                string `hl7:"TCD.43"`  // 10
}

// OBX - observation/result segment
type OBX struct {
	SetID                                        string `hl7:"OBX.1"`    // 1
	ValueType                                    string `hl7:"OBX.2"`    // 2
	ObservationIdentifier                        string `hl7:"OBX.3"`    // 3
	ObservationSubID                             string `hl7:"OBX.4"`    // 4
	ObservationValue                             string `hl7:"OBX.5"`    // 5
	Units                                        string `hl7:"OBX.6"`    // 6
	ReferencesRangeTargetRange                   string `hl7:"OBX.7.0"`  // 7.1
	ReferencesRangeTargetValue                   string `hl7:"OBX.7.1"`  // 7.2
	ReferencesRangeStandardDeviation             string `hl7:"OBX.7.2"`  // 7.3
	AbnormalFlags                                string `hl7:"OBX.8"`    // 8
	Probability                                  string `hl7:"OBX.9"`    // 9		not used
	NatureOfAbnormalTest                         string `hl7:"OBX.10"`   // 10		not used
	ObservationResultStatus                      string `hl7:"OBX.11"`   // 11
	DateLastObs                                  string `hl7:"OBX.12"`   // 12		not used
	UserDefAccessChecks                          string `hl7:"OBX.13"`   // 13
	ObservationTimestamp                         string `hl7:"OBX.14"`   // 14
	ProducerID1                                  string `hl7:"OBX.15.0"` // 15.1	not used
	ProducerID2                                  string `hl7:"OBX.15.1"` // 15.2	not used
	ProducerID3                                  string `hl7:"OBX.15.2"` // 15.3	not used
	ProducerIID                                  string `hl7:"OBX.15.3"` // 15.4
	OperatorIdentificationInstrumentUser         string `hl7:"OBX.16.0"` // 16.1
	OperatorIdentificationValidationUser         string `hl7:"OBX.16.1"` // 16.2
	ObservationMethod                            string `hl7:"OBX.17"`   // 17
	EquipmentInstanceIdentifier                  string `hl7:"OBX.18.0"` // 18.1
	EquipmentInstanceIdentifier1                 string `hl7:"OBX.18.1"` // 18.2	not used
	EquipmentInstanceIdentifier2                 string `hl7:"OBX.18.2"` // 18.3	not used
	EquipmentInstanceIdentifier3                 string `hl7:"OBX.18.3"` // 18.4	not used
	EquipmentInstanceIdentifierStandbyBottleFlag string `hl7:"OBX.18.4"` // 18.5
	AnalysisTimestamp                            string `hl7:"OBX.19"`   // 19
}

// EQU - equipment detail segment
type EQU struct {
	EquipmentIdentifier      string `hl7:"EQU.1.0"` // 1.1
	EquipmentIdentifierEvent string `hl7:"EQU.1.1"` // 1.2
	EventTimestamp           string `hl7:"EQU.2"`   // 2
	EquipmentState           string `hl7:"EQU.3"`   // 3		not used
	LocalRemoteControlState  string `hl7:"EQU.4"`   // 4		not used
	AlertLevel               string `hl7:"EQU.5"`   // 5		not used
}

// SID - substance identifier segment
type SID struct {
	AnalyteID                     string `hl7:"SID.1.0"` // 1.1
	AnalyteID1                    string `hl7:"SID.1.1"` // 1.1
	AnalyteID2                    string `hl7:"SID.1.2"` // 1.1
	AnalyteReagentTypeInformation string `hl7:"SID.1.3"` // 1.4
	ReagentLotID                  string `hl7:"SID.2"`   // 2
	BottleCountNumberKitID        string `hl7:"SID.3"`   // 3
}

// ZRI - reagent type information segment
type ZRI struct {
	ReagentTypeInformation string `hl7:"ZRI.1"` // 1
}

// SAC - specimen and container detail segment
type SAC struct {
	ExternalAccessionIdentifier        string `hl7:"SAC.1"`    // 1		not used
	AccessionIdentifier                string `hl7:"SAC.2"`    // 2		not used
	ContainerIdentifierSampleID        string `hl7:"SAC.3.0"`  // 3.1
	ContainerIdentifierControlCode     string `hl7:"SAC.3.1"`  // 3.2
	ContainerIdentifierLotNumber       string `hl7:"SAC.3.2"`  // 3.3
	ParentContainerIdentifier          string `hl7:"SAC.4"`    // 4
	EquipmentContainerIdentifier       string `hl7:"SAC.5"`    // 5		not used
	SampleSource                       string `hl7:"SAC.6"`    // 6
	RegistrationTimestamp              string `hl7:"SAC.7"`    // 7
	ContainerStatus                    string `hl7:"SAC.8"`    // 8
	CarrierType                        string `hl7:"SAC.9"`    // 9
	CarrierIdentifier                  string `hl7:"SAC.10"`   // 10
	PositionInCarrier                  string `hl7:"SAC.11"`   // 11
	TrayType                           string `hl7:"SAC.12"`   // 12	not used
	TrayIdentifier                     string `hl7:"SAC.13"`   // 13	not used
	PositionInTray                     string `hl7:"SAC.14"`   // 14	not used
	Location                           string `hl7:"SAC.15"`   // 15
	ContainerHeight                    string `hl7:"SAC.16"`   // 16	not used
	ContainerDiameter                  string `hl7:"SAC.17"`   // 17	not used
	BarrierDelta                       string `hl7:"SAC.18"`   // 18	not used
	BottomDelta                        string `hl7:"SAC.19"`   // 19	not used
	ContainerHeightDiameterUnits       string `hl7:"SAC.20"`   // 20	not used
	ContainerVolume                    string `hl7:"SAC.21"`   // 21	not used
	SampleVolume                       string `hl7:"SAC.22"`   // 22
	InitialSampleVolumeUpperLevelSerum string `hl7:"SAC.23.0"` // 2.1
	InitialSampleVolumeLowerLevelSerum string `hl7:"SAC.23.1"` // 2.2
	InitialSampleVolumeUpperLevelClot  string `hl7:"SAC.23.2"` // 2.3
	InitialSampleVolumeLowerLevelClot  string `hl7:"SAC.23.3"` // 2.4
	VolumeUnits                        string `hl7:"SAC.24"`   // 24
	SeparatorType                      string `hl7:"SAC.25"`   // 25	not used
	CapType                            string `hl7:"SAC.26"`   // 26	not used
	Additive                           string `hl7:"SAC.27"`   // 27	not used
	SampleSpecimenComponent            string `hl7:"SAC.28"`   // 28	not used
	DilutionFactor                     string `hl7:"SAC.29"`   // 23	not used
	Treatment                          string `hl7:"SAC.30"`   // 30	not used
	Temperature                        string `hl7:"SAC.31"`   // 31	not used
	HemolysisIndex                     string `hl7:"SAC.32"`   // 32	not used
	HemolysisIndexUnits                string `hl7:"SAC.33"`   // 33	not used
	LipemiaIndex                       string `hl7:"SAC.34"`   // 34	not used
	LipemiaIndexUnits                  string `hl7:"SAC.35"`   // 35	not used
	IcterusIndex                       string `hl7:"SAC.36"`   // 36	not used
	IcterusIndexUnits                  string `hl7:"SAC.37"`   // 37	not used
	FibrinIndex                        string `hl7:"SAC.38"`   // 38	not used

}

// QPD - query parameter definition segment
type QPD struct {
	MessageQueryName                   string `hl7:"QPD.1"`   // 1		not used
	SampleID                           string `hl7:"QPD.2"`   // 3		not used
	SelectionMode                      string `hl7:"QPD.4"`   // 4
	SampleVolume                       string `hl7:"QPD.5.0"` // 5.1
	SampleVolumeUnit                   string `hl7:"QPD.5.1"` // 5.2
	InitialSampleVolumeUpperLevelSerum string `hl7:"QPD.6.0"` // 6.1
	InitialSampleVolumeLowerLevelSerum string `hl7:"QPD.6.1"` // 6.2
	InitialSampleVolumeUpperLevelClot  string `hl7:"QPD.6.2"` // 6.3
	InitialSampleVolumeLowerLevelClot  string `hl7:"QPD.6.3"` // 6.4
}

// ZPD or PCD - patient code details segment
// This segment is only received by cobas IT middleware
type PCD struct {
	ExpectedBirthDate string `hl7:"PCD.1"` // 1
	MenstruationCycle string `hl7:"PCD.2"` // 2
	ActiveMedication  string `hl7:"PCD.3"` // 3		not used
	Species           string `hl7:"PCD.4"` // 4		not used
}
type ZPE struct {
	ExpectedBirthDate string `hl7:"ZPD.1"` // 1
	MenstruationCycle string `hl7:"ZPD.2"` // 2
	ActiveMedication  string `hl7:"ZPD.3"` // 3		not used
	Species           string `hl7:"ZPD.4"` // 4		not used
}

// PV1 - patient visit segment
// This segment is only received by cobas IT middleware
type PV1 struct {
	SetID                   string `hl7:"PV1.1.0"` // 1.1
	PatientClass            string `hl7:"PV1.1.1"` // 1.2		not used
	AssignedPatientLocation string `hl7:"PV1.1.2"` // 1.3
	AdmissionType           string `hl7:"PV1.1.3"` // 1.4		not used
	PreadmitNumber          string `hl7:"PV1.1.4"` // 1.5		not used
	PriorPatientLocation    string `hl7:"PV1.1.5"` // 1.6		not used
	AttendingDoctor         string `hl7:"PV1.1.6"` // 1.7		not used
	ReferringDoctor         string `hl7:"PV1.1.7"` // 1.8		not used
	ConsultingDoctor        string `hl7:"PV1.1.8"` // 1.9		not used
	HospitalService         string `hl7:"PV1.1.9"` // 1.10
}
