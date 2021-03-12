package enum_test

import (
	"testing"

	"github.com/neoxelox/enum"
)

type State string

type enumStates = struct {
	COMMITTED   State
	IN_PROGRESS State
	DONE        State
	enum.Enum
}

func TestNew(t *testing.T) {
	States := enum.New(&enumStates{
		COMMITTED:   "COMMITTED",
		IN_PROGRESS: "BLOCKED",
		DONE:        "DONE",
	}).(*enumStates)

	if States.COMMITTED != "COMMITTED" {
		t.Fatalf("Failed committed test")
	}
}
