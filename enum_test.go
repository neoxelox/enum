package enum_test

import (
	"reflect"
	"testing"

	"github.com/neoxelox/enum"
)

type LiteState string

type enumLiteStates = struct {
	COMMITTED   LiteState
	IN_PROGRESS LiteState
	DONE        LiteState
	enum.Enum
}

func TestNewLite(t *testing.T) {
	liteStates := enum.New(&enumLiteStates{
		COMMITTED:   "COMMITTED",
		IN_PROGRESS: "BLOCKED",
		DONE:        "DONE",
	}).(*enumLiteStates)

	if reflect.TypeOf(liteStates.COMMITTED).Kind() != reflect.String {
		t.Fatalf("Type of Lite enum is not a primitive string")
	}

	if liteStates.COMMITTED != "COMMITTED" {
		t.Fatalf("Lite enum is not a comparable with an string literal")
	}
}

func TestIsLite(t *testing.T) {
	liteStates := enum.New(&enumLiteStates{
		COMMITTED:   "COMMITTED",
		IN_PROGRESS: "BLOCKED",
		DONE:        "DONE",
	}).(*enumLiteStates)

	if !liteStates.Is("IN_PROGRESS") {
		t.Fatalf("IN_PROGRESS should be an alias of Lite enum")
	}

	if liteStates.Is("BLOCKED") {
		t.Fatalf("BLOCKED should not be an alias of Lite enum")
	}
}

func TestAliasesLite(t *testing.T) {
	liteStates := enum.New(&enumLiteStates{
		COMMITTED:   "COMMITTED",
		IN_PROGRESS: "BLOCKED",
		DONE:        "DONE",
	}).(*enumLiteStates)

	aliases := liteStates.Aliases()
	if len(aliases) != 3 {
		t.Fatalf("Aliases of Lite enum should be 3, instead got: %v", len(aliases))
	}

	expectedAliases := []string{"COMMITTED", "IN_PROGRESS", "DONE"}
	for i := range aliases {
		if aliases[i] != expectedAliases[i] {
			t.Fatalf("Expected alias: %v, instead got: %v, in Lite enum", expectedAliases[i], aliases[i])
		}
	}
}

func TestInLite(t *testing.T) {
	liteStates := enum.New(&enumLiteStates{
		COMMITTED:   "COMMITTED",
		IN_PROGRESS: "BLOCKED",
		DONE:        "DONE",
	}).(*enumLiteStates)

	if !liteStates.In(LiteState("BLOCKED")) {
		t.Fatalf("BLOCKED should be a value of Lite enum")
	}

	if liteStates.In(LiteState("IN_PROGRESS")) {
		t.Fatalf("IN_PROGRESS should not be a value of Lite enum")
	}
}

func TestValuesLite(t *testing.T) {
	liteStates := enum.New(&enumLiteStates{
		COMMITTED:   "COMMITTED",
		IN_PROGRESS: "BLOCKED",
		DONE:        "DONE",
	}).(*enumLiteStates)

	values := liteStates.Values()
	if len(values) != 3 {
		t.Fatalf("Values of Lite enum should be 3, instead got: %v", len(values))
	}

	expectedValues := []LiteState{LiteState("COMMITTED"), LiteState("BLOCKED"), LiteState("DONE")}
	for i := range values {
		if alias := (values[i]).(LiteState); alias != expectedValues[i] {
			t.Fatalf("Expected alias: %v, instead got: %v, in Lite enum", expectedValues[i], alias)
		}
	}
}

type StrongState struct{ string }

type enumStrongStates = struct {
	enum.Enum
	COMMITTED   StrongState
	IN_PROGRESS StrongState
	DONE        StrongState
}

func TestNewStrong(t *testing.T) {
	strongStates := enum.New(&enumStrongStates{
		COMMITTED:   StrongState{"COMMITTED"},
		IN_PROGRESS: StrongState{"BLOCKED"},
		DONE:        StrongState{"DONE"},
	}).(*enumStrongStates)

	if reflect.TypeOf(strongStates.COMMITTED).Kind() != reflect.Struct {
		t.Fatalf("Type of Strong enum is not a primitive struct")
	}
}

func TestIsStrong(t *testing.T) {
	strongStates := enum.New(&enumStrongStates{
		COMMITTED:   StrongState{"COMMITTED"},
		IN_PROGRESS: StrongState{"BLOCKED"},
		DONE:        StrongState{"DONE"},
	}).(*enumStrongStates)

	if !strongStates.Is("IN_PROGRESS") {
		t.Fatalf("IN_PROGRESS should be an alias of Strong enum")
	}

	if strongStates.Is("BLOCKED") {
		t.Fatalf("BLOCKED should not be an alias of Strong enum")
	}
}

func TestAliasesStrong(t *testing.T) {
	strongStates := enum.New(&enumStrongStates{
		COMMITTED:   StrongState{"COMMITTED"},
		IN_PROGRESS: StrongState{"BLOCKED"},
		DONE:        StrongState{"DONE"},
	}).(*enumStrongStates)

	aliases := strongStates.Aliases()
	if len(aliases) != 3 {
		t.Fatalf("Aliases of Strong enum should be 3, instead got: %v", len(aliases))
	}

	expectedAliases := []string{"COMMITTED", "IN_PROGRESS", "DONE"}
	for i := range aliases {
		if aliases[i] != expectedAliases[i] {
			t.Fatalf("Expected alias: %v, instead got: %v, in Strong enum", expectedAliases[i], aliases[i])
		}
	}
}

func TestInStrong(t *testing.T) {
	strongStates := enum.New(&enumStrongStates{
		COMMITTED:   StrongState{"COMMITTED"},
		IN_PROGRESS: StrongState{"BLOCKED"},
		DONE:        StrongState{"DONE"},
	}).(*enumStrongStates)

	if !strongStates.In(StrongState{"BLOCKED"}) {
		t.Fatalf("BLOCKED should be a value of Strong enum")
	}

	if strongStates.In(StrongState{"IN_PROGRESS"}) {
		t.Fatalf("IN_PROGRESS should not be a value of Strong enum")
	}
}

func TestValuesStrong(t *testing.T) {
	strongStates := enum.New(&enumStrongStates{
		COMMITTED:   StrongState{"COMMITTED"},
		IN_PROGRESS: StrongState{"BLOCKED"},
		DONE:        StrongState{"DONE"},
	}).(*enumStrongStates)

	values := strongStates.Values()
	if len(values) != 3 {
		t.Fatalf("Values of Strong enum should be 3, instead got: %v", len(values))
	}

	expectedValues := []StrongState{{"COMMITTED"}, {"BLOCKED"}, {"DONE"}}
	for i := range values {
		if alias := (values[i]).(StrongState); alias != expectedValues[i] {
			t.Fatalf("Expected alias: %v, instead got: %v, in Strong enum", expectedValues[i], alias)
		}
	}
}
