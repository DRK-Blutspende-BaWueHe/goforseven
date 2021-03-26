package hl7model

type fieldtag struct {
	Identifier  string
	Segmentname string
	Position    []string
}

type Delimeter struct {
	Parentlevel string
	Childlevels string
	CrLf        string
}

var delimeterwork Delimeter
