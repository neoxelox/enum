// Package enum implements an enumerator and functions to work with.
package enum

import (
	"reflect"
)

type enumerator interface {
	set(enum enumerator)
}

// Enum describes an Enumerator.
type Enum struct {
	aliases []string
	values  []interface{}
}

// New creates a new Enum.
func New(enum enumerator) enumerator {
	enum.set(enum)

	return enum
}

func (e *Enum) set(enum enumerator) {
	enumRef := reflect.ValueOf(enum).Elem()
	numFields := enumRef.NumField() - 1

	e.aliases = make([]string, numFields)
	e.values = make([]interface{}, numFields)

	for i := 0; i < numFields; i++ {
		e.aliases[i] = enumRef.Type().Field(i).Name
		e.values[i] = enumRef.Field(i).Interface()
	}
}

// Is checks whether an alias is present in the enum or not.
func (e *Enum) Is(alias string) bool {
	for _, setAlias := range e.aliases {
		if setAlias == alias {
			return true
		}
	}

	return false
}

// Aliases returns all the aliases of the enum.
func (e *Enum) Aliases() []string {
	return e.aliases
}

// In checks whether a value is present in the enum or not.
func (e *Enum) In(value interface{}) bool {
	for _, setValue := range e.values {
		if setValue == value {
			return true
		}
	}

	return false
}

// Values returns all the values of the enum.
func (e *Enum) Values() []interface{} {
	return e.values
}
