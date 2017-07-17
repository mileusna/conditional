package conditional_test

import (
	"testing"

	"github.com/mileusna/conditional"
)

func TestConditional(t *testing.T) {

	if conditional.Int(true, 5, 4) != 5 {
		t.Fatal("Error Int")
	}

	if conditional.Int(false, 5, 4) != 4 {
		t.Fatal("Error Int")
	}

	if conditional.String(true, "ok", "not ok") != "ok" {
		t.Fatal("Error string")
	}

	if conditional.String(false, "ok", "not ok") != "not ok" {
		t.Fatal("Error string")
	}

	if conditional.UInt(true, 5, 4) != 5 {
		t.Fatal("Error UInt")
	}

	if conditional.UInt(false, 5, 4) != 4 {
		t.Fatal("Error UInt")
	}

}
