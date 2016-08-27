package parser

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"os"
)

func TestJSONParser(t *testing.T) {
	f, _ := os.Open("../testdata/test.json")
	fInvalid, _ := os.Open("../testdata/invalid.json")
	Convey("Should parse a json", t, func() {
		clients, err := ToJSON(f)
		So(err, ShouldBeNil)
		So(len(clients), ShouldEqual, 1)
		So(clients[0].Nome, ShouldEqual, "Felipe")
		So(clients[0].Sexo, ShouldEqual, "M")
		So(clients[0].Idade, ShouldEqual, 23)
		So(clients[0].Email, ShouldEqual, "Test@teste.com.br")
		So(clients[0].Outros["abobrinha"], ShouldEqual, "Verde")
		So(clients[0].Outros["banana"], ShouldEqual, "Amarela")
	})
	Convey("Should give an error if json invalid", t, func() {
		_, err := ToJSON(fInvalid)
		So(err, ShouldNotBeNil)
	})
}

func TestCSVParser(t *testing.T) {
	f, _ := os.Open("../testdata/test.csv")
	fInvalid, _ := os.Open("../testdata/invalid.csv")
	Convey("Should parse a csv", t, func() {
		clients, err := ToCSV(f)
		So(err, ShouldBeNil)
		So(len(clients), ShouldEqual, 1)
		So(clients[0].Nome, ShouldEqual, "Felipe")
		So(clients[0].Sexo, ShouldEqual, "M")
		So(clients[0].Idade, ShouldEqual, 23)
		So(clients[0].Email, ShouldEqual, "Test@teste.com.br")
		So(clients[0].Outros["abobrinha"], ShouldEqual, "Verde")
		So(clients[0].Outros["banana"], ShouldEqual, "Amarela")
	})
	Convey("Should give an error if csv invalid", t, func() {
		_, err := ToCSV(fInvalid)
		So(err, ShouldNotBeNil)
	})
}
