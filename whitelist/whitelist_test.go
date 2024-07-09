package whitelist

import "testing"

func TestContracts(t *testing.T) {
	if Contracts == "" {
		t.Error("Contracts variable is empty")
	}
}
