package http

import (
	"errors"
	"github.com/felipeweb/nuveo/models"
	"github.com/felipeweb/nuveo/parser"
	"net/http"
	"strings"
)

const (
	CSV  = "CSV"
	JSON = "JSON"
)

func ProcessFile(url string) ([]models.Client, error) {
	var clients []models.Client
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	contentType, err := checkContentTypeIsValid(resp.Header["Content-Type"])

	switch contentType {
	case JSON:
		clients, err = parser.ToJSON(resp.Body)
	case CSV:
		clients, err = parser.ToCSV(resp.Body)
	}

	return clients, err
}

func checkContentTypeIsValid(contentTypes []string) (string, error) {
	for _, contentType := range contentTypes {
		ct := strings.ToLower(contentType)
		if strings.Contains(ct, "/csv") {
			return CSV, nil
		} else if strings.Contains(ct, "/json") {
			return JSON, nil
		}
	}
	return "", errors.New("The file is not a CSV or JSON file")
}
