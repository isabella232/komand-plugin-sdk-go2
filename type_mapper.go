package main

import (
	"fmt"
	"strings"
)

// TypeMapper handles renaming types from the spec when they don't line up with go types by name
type TypeMapper struct {
	CustomTypes map[string]string
}

// NewTypeMapper returns a new TypeMapper
func NewTypeMapper(s *PluginSpec) *TypeMapper { // TO FIGURE OUT object and bytes need some work, look at generate/conversion
	t := &TypeMapper{
		CustomTypes: map[string]string{
			"boolean":   "bool",
			"object":    "map[string]interface{}",
			"integer":   "int",
			"bytes":     "[]byte",
			"date":      "time.Time",
			"date-time": "time.Time",
			"number":    "float64",
			"float":     "float64",
		},
	}

	for name := range s.RawTypes {
		goName := UpperCamelCase(name)
		tName := "types." + goName
		/*var isArr bool
		// If it's an array tpye we nee
		if strings.HasPrefix(nsme, "[]") {
			name = name[2:]
			isArr = true
			tName =
		}*/

		t.CustomTypes[name] = tName
	}
	return t
}

// SpecTypeToGoType maps spec types to go types
func (t *TypeMapper) SpecTypeToGoType(specType string) string {
	var isArr bool
	strippedType := specType
	if strings.HasPrefix(specType, "[]") {
		strippedType = specType[2:]
		isArr = true
	}
	foundType, ok := t.CustomTypes[strippedType]
	if !ok {
		// We don't have a remap for it, so use the original
		return specType
	}
	if isArr {
		fmt.Println(foundType)
		foundType = "[]" + foundType
	}
	return foundType
}
