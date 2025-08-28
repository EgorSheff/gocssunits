package gocssunits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	fmt.Println(ParseFontSize("14.2px"))
	fmt.Println(ParseFontSize("14"))
	fmt.Println(ParseFontSize("large"))
	fmt.Println(ParseFontSize("invalid"))
}

func TestJSON(t *testing.T) {
	fz1, _ := ParseFontSize("14.66em")

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(fz1); err != nil {
		t.Fatal(err)
	}

	var fz2 FontSize
	if err := json.NewDecoder(&buf).Decode(&fz2); err != nil {
		t.Fatal(err)
	}

	if fz1.Unit != fz2.Unit || fz1.Value != fz2.Value {
		t.Fatal("not matched", fz1, fz2)
	}
}
