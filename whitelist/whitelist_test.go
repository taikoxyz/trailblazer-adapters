package whitelist_test

import (
	"testing"

	"github.com/taikoxyz/trailblazer-adapters/whitelist"
)

// TestProtocols checks if the JSON data is correctly unmarshaled into Protocol structs
func TestProtocols(t *testing.T) {
	protocols, err := whitelist.Protocols()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(protocols) == 0 {
		t.Fatalf("Expected some protocols, got none")
	}

	t.Run("mandatory fields are correctly parsed", func(t *testing.T) {
		for _, protocol := range protocols {
			if protocol.Name == "" {
				t.Errorf("Expected protocol name, got empty string")
			}

			if protocol.Slug == "" {
				t.Errorf("Expected protocol slug, got empty string")
			}

			if len(protocol.Contracts) == 0 {
				t.Errorf("Expected some contracts for protocol %s, got none", protocol.Name)
			}

			for _, contract := range protocol.Contracts {
				if len(contract) != 42 || contract[:2] != "0x" {
					t.Errorf("Invalid contract address format: %s", contract)
				}
			}
		}
	})
}
