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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	contentType, err := validContentType(resp.Header["Content-Type"])
	if err != nil {
		return nil, err
	}
	if contentType == JSON {
		return parser.ToJSON(resp.Body)
	} else {
		return parser.ToCSV(resp.Body)
	}
}

func validContentType(contentTypes []string) (string, error) {
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
