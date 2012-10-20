// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysqlctl

import (
	"testing"
)

func compareArrays(t *testing.T, actual, expected []string) {
	equal := false
	if len(actual) == len(expected) {
		equal = true
		for i, val := range actual {
			if val != expected[i] {
				equal = false
				break
			}
		}
	}

	if !equal {
		t.Logf("Expected: %v", expected)
		t.Logf("Actual: %v", actual)
		t.Fail()
	}
}

func TestSchemaDiff(t *testing.T) {
	sd1 := &SchemaDefinition{TableDefinitions: make([]TableDefinition, 2)}
	sd1.TableDefinitions[0].Name = "table1"
	sd1.TableDefinitions[0].Schema = "schema1"
	sd1.TableDefinitions[1].Name = "table2"
	sd1.TableDefinitions[1].Schema = "schema2"
	compareArrays(t, sd1.diffSchema("sd1", "sd2", sd1), []string{})

	sd2 := &SchemaDefinition{TableDefinitions: make([]TableDefinition, 0, 2)}
	compareArrays(t, sd2.diffSchema("sd1", "sd2", sd2), []string{})
	compareArrays(t, sd1.diffSchema("sd1", "sd2", sd2), []string{"sd1 has an extra table named table1", "sd1 has an extra table named table2"})
	compareArrays(t, sd2.diffSchema("sd2", "sd1", sd1), []string{"sd1 has an extra table named table1", "sd1 has an extra table named table2"})

	sd2.TableDefinitions = append(sd2.TableDefinitions, TableDefinition{Name: "table1", Schema: "schema1"})
	compareArrays(t, sd1.diffSchema("sd1", "sd2", sd2), []string{"sd1 has an extra table named table2"})

	sd2.TableDefinitions = append(sd2.TableDefinitions, TableDefinition{Name: "table2", Schema: "schema3"})
	compareArrays(t, sd1.diffSchema("sd1", "sd2", sd2), []string{"sd1 and sd2 disagree on schema for table table2"})
}