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
	pass := genPass(vc.site, "73f", 1234)
	if pass != vc.password {
		t.Fatalf("%s is not equal to %s", pass, vc.password)
	}

}
