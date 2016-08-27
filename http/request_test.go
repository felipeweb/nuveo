package http

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContentType(t *testing.T) {
	invalid := []string{"ovdovdvodv"}
	validJson := []string{"application/json"}
	validCsv := []string{"application/csv"}

	Convey("Should verify json content type", t, func() {
		contentType, err := validContentType(validJson)
		So(err, ShouldBeNil)
		So(contentType, ShouldEqual, JSON)
	})
	Convey("Should verify csv content type", t, func() {
		contentType, err := validContentType(validCsv)
		So(err, ShouldBeNil)
		So(contentType, ShouldEqual, CSV)
	})
	Convey("Should verify invalid content type", t, func() {
		_, err := validContentType(invalid)
		So(err, ShouldNotBeNil)
	})
}

func TestProcessFile(t *testing.T) {
	csvURL := "http://www.uk-postcodes.com/postcode/BD72TA.csv"
	jsonURL := "http://www.uk-postcodes.com/postcode/BD72TA.json"
	html := "https://www.felipeweb.com.br"

	Convey("Should parse json file", t, func() {
		_, err := ProcessFile(jsonURL)
		So(err, ShouldBeNil)
	})
	Convey("Should parse csv file", t, func() {
		_, err := ProcessFile(csvURL)
		So(err, ShouldBeNil)
	})
	Convey("Should parse html file", t, func() {
		_, err := ProcessFile(html)
		So(err, ShouldNotBeNil)
	})
}