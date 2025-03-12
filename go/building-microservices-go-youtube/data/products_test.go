package data

import "testing"

func TestDataValidation(t *testing.T) {
	p := Product{Name: "hello", SKU: "abc-abc-abc"}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
