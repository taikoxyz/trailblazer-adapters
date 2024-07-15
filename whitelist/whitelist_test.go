package whitelist

import (
	"encoding/json"
	"testing"
)

// Test data
var contractsData = Contracts

// Helper function to unmarshal JSON data
func unmarshalContracts(data string) ([]Protocol, error) {
	var protocols []Protocol
	err := json.Unmarshal([]byte(data), &protocols)
	return protocols, err
}

// TestUnmarshalContracts checks if the JSON data is correctly unmarshaled into Protocol structs
func TestUnmarshalContracts(t *testing.T) {
	protocols, err := unmarshalContracts(contractsData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(protocols) == 0 {
		t.Fatalf("Expected some protocols, got none")
	}
}

// TestProtocolFields checks if the essential fields of the Protocol struct are correctly parsed
func TestProtocolFields(t *testing.T) {
	protocols, err := unmarshalContracts(contractsData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for _, protocol := range protocols {
		if protocol.Name == "" {
			t.Errorf("Expected protocol name, got empty string")
		}

		if protocol.Slug == "" {
			t.Errorf("Expected protocol slug, got empty string")
		}

		if len(protocol.Contracts) == 0 {
			t.Errorf("Expected some contracts for protocol %v, got none", protocol.Name)
		}
	}
}

// TestContractsFormat checks if the contract addresses are in the correct format
func TestContractsFormat(t *testing.T) {
	protocols, err := unmarshalContracts(contractsData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for _, protocol := range protocols {
		for _, contract := range protocol.Contracts {
			if len(contract) != 42 || contract[:2] != "0x" {
				t.Errorf("Invalid contract address format: %v", contract)
			}
		}
	}
}
