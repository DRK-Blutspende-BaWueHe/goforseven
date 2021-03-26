package hl7model

import (
	"errors"
	"strings"
)

func getFieldTag(tag string) (fieldtag, error) {
	fieldtag := fieldtag{}

	tags := strings.Split(tag, ":")
	fieldtag.Identifier = tags[0]
	if len(tags) > 1 {
		positionvalues := strings.Split(tags[1], ".")
		fieldtag.Segmentname = positionvalues[0]
		for i := 1; i < len(positionvalues); i++ {
			fieldtag.Position = append(fieldtag.Position, positionvalues[i])
		}
	} else {
		err := errors.New("Invalid tag-format")
		return fieldtag, err
	}

	return fieldtag, nil
}
