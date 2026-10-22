package views

import (
	"encoding/json"
	"encoding/xml"
)

const (
	jsonType = "application/json"
	xmlType  = "application/xml"
)

type formatedResult struct {
	Body        []byte
	ContentType string
}

func FormatJson(data any) (formatedResult, error) {
	formatedBody, err := json.Marshal(data)

	if err != nil {
		return formatedResult{}, err
	}

	return formatedResult{
		Body:        formatedBody,
		ContentType: jsonType,
	}, nil
}

func FormatXml(data any) (formatedResult, error) {
	formatedBody, err := xml.Marshal(data)

	if err != nil {
		return formatedResult{}, err
	}

	return formatedResult{
		Body:        formatedBody,
		ContentType: xmlType,
	}, nil
}
