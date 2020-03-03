package generation

import (
	"testing"
)

type validCombo struct {
	site     string
	password string
}

func TestGeneration(t *testing.T) {
	vc := validCombo{
		site:     "186806",
		password: "04d231a5e",
	}

	_, err := genPass(vc.site, "73f", 65536)
	if err == nil {
		t.Fatal("expected password > 65535 error!")
	}
	_, _, err = GenerateCredentials("73f", 10000000, 1234)
	if err == nil {
		t.Fatal("expected site > 999999 error!")
	}
	pass, err := genPass(vc.site, "73f", 1234)
	if err != nil {
		t.Fatal(err)
	}
	if pass != vc.password {
		t.Fatalf("%s is not equal to %s!", pass, vc.password)
	}
	_, _, err = GenerateCredentials("73f", 0, 0)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}
