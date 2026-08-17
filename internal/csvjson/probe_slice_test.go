package csvjson

import "testing"

func TestProbeEmptyColumnsRejectRows(t *testing.T) {
	if _, err := JSONToCSV([]string{}, [][]any{{"orphan"}}, ','); err == nil {
		t.Fatal("rows must be rejected when columns is empty")
	}
}
